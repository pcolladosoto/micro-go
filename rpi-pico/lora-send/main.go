package main

import (
	"time"

	"machine"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/sx127x"
)

const (
	START_DELAY int    = 10
	MSG_CONST   string = "pico,"
	TX_TIMEOUT  uint32 = 500
)

func main() {
	machine.InitSerial()

	for i := 0; i < START_DELAY; i++ {
		println("starting in", START_DELAY-i, "seconds...")
		time.Sleep(1 * time.Second)
	}

	spiPort := machine.SPI0
	if err := spiPort.Configure(machine.SPIConfig{
		Frequency: 5_000_000,
		LSBFirst:  false,
		Mode:      machine.Mode0,
		SCK:       machine.SPI0_SCK_PIN,
		SDI:       machine.SPI0_SDI_PIN,
		SDO:       machine.SPI0_SDO_PIN,
	}); err != nil {
		println("error configuring the SPI0 port:", err)
	}

	println("detected SPI0 baudrate:", spiPort.GetBaudRate())
	print("SPI registers:")
	spiPort.PrintRegs()

	rstPin := machine.GP20
	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	radio := sx127x.New(*spiPort, rstPin)
	radioController := sx127x.NewRadioControl(machine.GP17, machine.GP21, machine.NoPin)
	if err := radio.SetRadioController(radioController); err != nil {
		println("error setting the radio controller:", err)
	}

	radio.Reset()

	println("radio registers ")
	radio.PrintRegisters(false)

	println("detected radio version:", radio.GetVersion())

	radio.LoraConfig(lora.Config{
		Freq:           868 * machine.MHz,
		Cr:             lora.CodingRate4_5,
		Sf:             lora.SpreadingFactor12,
		Bw:             lora.Bandwidth_125_0,
		Preamble:       8,
		Crc:            lora.CRCOn,
		LoraTxPowerDBm: 13,
	})

	println("radio registers after configuration:")
	radio.PrintRegisters(false)

	onboardLed := machine.LED

	// We need to initialise it as an output
	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// adc := machine.ADC{Pin: machine.ADC0}
	// machine.InitADC()
	// adc.Configure(machine.ADCConfig{})

	radioMsg := make([]byte, len(MSG_CONST)+2)
	println("copied ", copy(radioMsg[:len(MSG_CONST)], MSG_CONST), " bytes to msg slice")

	var dummyCnt uint16 = 0
	for {
		onboardLed.High()
		// adcVal := (adc.Get() >> 6) & 0x3FF

		// println("current ADC value: ", adcVal)

		radioMsg[len(radioMsg)-2] = byte((dummyCnt >> 8) & 0xFF)
		radioMsg[len(radioMsg)-1] = byte(dummyCnt & 0xFF)

		dummyCnt++

		for _, b := range radioMsg {
			print(b, " ")
		}
		println()

		println("sending message...")
		if err := radio.Tx(radioMsg, TX_TIMEOUT); err != nil {
			println("error sending data:", err)
		}

		println("done! waiting for the next one...")
		// println("detected radio version:", radio.GetVersion())

		time.Sleep(time.Second * 5)

		onboardLed.Low()

		time.Sleep(1 * time.Second)
	}
}
