package main

import (
	"machine"
	"time"
)

func main() {
	// Get a hold of the user LED (i.e. the one on the board)
	userLED := machine.LED

	// Configure the LED as an output
	userLED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Instantiate the ADC so that we can use it later on.
	adc := machine.ADC{Pin: machine.ADC0}

	// Initialize the ADC once it's configured
	machine.InitADC()

	// Configure the ADC. Passing a zeroed strcuct assumes the machine's default configuration.
	// Check https://tinygo.org/docs/reference/microcontrollers/machine/arduino/#type-adcconfig
	// for more information on what can be configured.
	adc.Configure(machine.ADCConfig{})

	// Initialise a counter to toggle the LED every ten reads
	readCnt := 1

	// Time to run!
	for {
		// Get a data reading
		adcVal := adc.Get()

		// Print the 10 MSB bits to the serial port: the six lower ones are to be ignored!
		println("Current ADC value: ", (adcVal >> 6) & 0x3FF)

		// Should we turn the LED on?
		if readCnt % 10 == 0 {
			userLED.High()
		} else if readCnt % 10 == 1 {
			userLED.Low()
		}

		// Keep on counting
		readCnt++
		
		// Wait a second before reading the next value
		time.Sleep(time.Second * 1)
	}
}
