package main

import (
	"time"

	"machine"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/sx127x"
)

const (
	START_DELAY_S uint8  = 10
	RX_TIMEOUT_MS uint32 = 5000
	LOOP_DELAY_S  uint8  = 1

	SPI_SCK_PIN machine.Pin = machine.SPI0_SCK_PIN
	SPI_SDI_PIN machine.Pin = machine.SPI0_SDI_PIN
	SPI_SDO_PIN machine.Pin = machine.SPI0_SDO_PIN
	RESET_PIN   machine.Pin = machine.GP20

	SPI_SCK_FREQ uint32 = 5 * machine.MHz
	SPI_MSB      bool   = true
	SPI_MODE     uint8  = machine.Mode0

	LORA_FREQ             uint32 = 868 * machine.MHz
	LORA_CODING_RATE      uint8  = lora.CodingRate4_5
	LORA_SPREADING_FACTOR uint8  = lora.SpreadingFactor7
	LORA_BANDWIDTH        uint8  = lora.Bandwidth_125_0
	LORA_PREAMBLE         uint16 = 8
	LORA_CRC              uint8  = lora.CRCOn
	LORA_TX_POWER_DBM     int8   = 13
)

var SPI_PORT *machine.SPI = machine.SPI0

func main() {
	machine.InitSerial()

	for i := uint8(0); i < START_DELAY_S; i++ {
		println("starting in", START_DELAY_S-i, "seconds...")
		time.Sleep(1 * time.Second)
	}

	spiPort := SPI_PORT
	if err := spiPort.Configure(machine.SPIConfig{
		Frequency: SPI_SCK_FREQ,
		LSBFirst:  !SPI_MSB,
		Mode:      SPI_MODE,
		SCK:       SPI_SCK_PIN,
		SDI:       SPI_SDI_PIN,
		SDO:       SPI_SDO_PIN,
	}); err != nil {
		println("error configuring the SPI0 port:", err)
	}

	println("detected SPI0 baudrate:", spiPort.GetBaudRate())

	rstPin := RESET_PIN
	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	radio := sx127x.New(*spiPort, rstPin)
	radioController := sx127x.NewRadioControl(machine.GP17, machine.GP21, machine.NoPin)
	if err := radio.SetRadioController(radioController); err != nil {
		println("error setting the radio controller:", err)
	}

	println("resetting the radio...")
	radio.Reset()

	println("checking the SPI bus...")
	if !radio.DetectDevice() {
		println("wrong radio version detected:", radio.GetVersion())
	}

	println("setting the LoRa parameters...")
	radio.LoraConfig(lora.Config{
		Freq:           LORA_FREQ,
		Cr:             LORA_CODING_RATE,
		Sf:             LORA_SPREADING_FACTOR,
		Bw:             LORA_BANDWIDTH,
		Preamble:       LORA_PREAMBLE,
		Crc:            LORA_CRC,
		LoraTxPowerDBm: LORA_TX_POWER_DBM,
	})

	onboardLed := machine.LED
	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	var msgCount uint16 = 0
	for {
		onboardLed.High()

		recvMsg, err := radio.Rx(RX_TIMEOUT_MS)
		if err != nil {
			println("error on reception:", err)
			continue
		} else if recvMsg == nil {
			// We had a timeout...
			continue
		}

		msgCount++

		println("received", msgCount, "messages. The new one is:")

		for _, b := range recvMsg[4:] {
			print(" ", b)
		}
		println("")

		onboardLed.Low()

		time.Sleep(time.Duration(LOOP_DELAY_S) * time.Second)
	}
}
