package main

import (
	"time"

	"machine"

	"github.com/ulbios/lora/sx1276-driver/arduino"
)

const (
	START_DELAY int    = 3
	MSG_CONST   string = "arduino,"
)

func main() {
	radioOpts := arduino.DefaultOpts

	for i := 0; i < START_DELAY; i++ {
		println("starting in", START_DELAY-i, "seconds...")
		time.Sleep(1 * time.Second)
	}

	radio, err := arduino.New(&radioOpts)
	if err != nil {
		println("error creating the radio", err)
	}

	adc := machine.ADC{Pin: machine.ADC0}
	machine.InitADC()
	adc.Configure(machine.ADCConfig{})

	radioMsg := make([]byte, len(MSG_CONST)+2)
	println("copied ", copy(radioMsg[:len(MSG_CONST)], MSG_CONST), " bytes to msg slice")

	for {
		adcVal := (adc.Get() >> 6) & 0x3FF

		println("current ADC value: ", adcVal)

		radioMsg[len(radioMsg)-2] = byte((adcVal >> 8) & 0xFF)
		radioMsg[len(radioMsg)-1] = byte(adcVal & 0xFF)

		for _, b := range radioMsg {
			print(b, " ")
		}
		println()

		if err := radio.Send(radioMsg); err != nil {
			println("error sending data", err.Error())
		}

		time.Sleep(time.Second * 5)
	}
}
