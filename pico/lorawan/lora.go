package main

import (
	"errors"
	"machine"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/sx127x"
)

func setupLoRa(rstPin machine.Pin, csPin machine.Pin, irqPin machine.Pin, spiBaudRate uint32, loraConf lora.Config) (*sx127x.Device, error) {
	spiPort := machine.SPI0
	if err := spiPort.Configure(machine.SPIConfig{
		Frequency: spiBaudRate,
		LSBFirst:  false,
		Mode:      machine.Mode0,
		SCK:       machine.SPI0_SCK_PIN,
		SDI:       machine.SPI0_SDI_PIN,
		SDO:       machine.SPI0_SDO_PIN,
	}); err != nil {
		return nil, err
	}

	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	radio := sx127x.New(*spiPort, rstPin)
	radioController := sx127x.NewRadioControl(csPin, irqPin, machine.NoPin)
	if err := radio.SetRadioController(radioController); err != nil {
		return nil, err
	}

	radio.Reset()

	if !radio.DetectDevice() {
		return nil, errors.New("detected a wrong radio version...")
	}

	radio.LoraConfig(loraConf)

	return radio, nil
}
