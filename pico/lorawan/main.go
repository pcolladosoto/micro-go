package main

// Please note the following is largely based on https://github.com/tinygo-org/drivers/blob/release/examples/lora/lorawan/basic-demo/main.go

import (
	"errors"
	"time"

	"machine"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/lora/lorawan"
	"tinygo.org/x/drivers/lora/lorawan/region"
)

const (
	START_DELAY int    = 10
	MSG_PREFIX  string = "e-raticDesperado:"

	SPI_BAUDRATE uint32 = 5 * machine.MHz

	RST_PIN machine.Pin = machine.GP20
	CS_PIN  machine.Pin = machine.GP17
	IRQ_PIN machine.Pin = machine.GP21

	LORA_FREQ   uint32 = 868 * machine.MHz
	LORA_CR     uint8  = lora.CodingRate4_5
	LORA_SF     uint8  = lora.SpreadingFactor12
	LORA_BW     uint8  = lora.Bandwidth_125_0
	LORA_PRE    uint16 = 8
	LORA_CRC    uint8  = lora.CRCOn
	LORA_TX_POW int8   = 13

	LORAWAN_JOIN_TIMEOUT    time.Duration = 180 * time.Second
	LORAWAN_RECONNECT_DELAY time.Duration = 15 * time.Second
	LORAWAN_UPLINK_DELAY    time.Duration = 60 * time.Second
)

var (
	APP_EUI []uint8 = []uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	DEV_EUI []uint8 = []uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	APP_KEY []uint8 = []uint8{0x12, 0x22, 0xA3, 0xFF, 0x0C, 0x7B, 0x76, 0x7B, 0x8F, 0xD3, 0x12, 0x4F, 0xCE, 0x7A, 0x32, 0x16}
)

func main() {
	for i := 0; i < START_DELAY; i++ {
		println("starting in", START_DELAY-i, "seconds...")
		time.Sleep(1 * time.Second)
	}

	radio, err := setupLoRa(RST_PIN, CS_PIN, IRQ_PIN, SPI_BAUDRATE, lora.Config{
		Freq:           LORA_FREQ,
		Cr:             LORA_CR,
		Sf:             LORA_SF,
		Bw:             LORA_BW,
		Preamble:       LORA_PRE,
		Crc:            LORA_CRC,
		LoraTxPowerDBm: LORA_TX_POW,
	})
	if err != nil {
		println("error instantiating the radio:", err)
		for {
		}
	}

	onboardLed := machine.LED
	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Required for LoraWan operations
	session := &lorawan.Session{}
	otaa := &lorawan.Otaa{}

	// Connect the lorawan with the Lora Radio device.
	lorawan.UseRadio(radio)

	// Set up the regional parameters
	lorawan.UseRegionSettings(region.EU868())

	// Set up the keys
	otaa.SetAppEUI(APP_EUI)
	otaa.SetDevEUI(DEV_EUI)
	otaa.SetAppKey(APP_KEY)

	// Use the public syncword (i.e. join public networks)
	lorawan.SetPublicNetwork(true)

	if err := connectLoRaWAN(otaa, session); err != nil {
		println("error joining the LoRaWAN network:", err)
		for {
		}
	}

	payload := make([]byte, len(MSG_PREFIX)+2)
	copy(payload[:len(MSG_PREFIX)], MSG_PREFIX)

	var frameCnt uint16 = 0
	for {
		onboardLed.High()

		payload[len(payload)-2] = byte((frameCnt >> 8) & 0xFF)
		payload[len(payload)-1] = byte(frameCnt & 0xFF)

		frameCnt++

		println("sending message...")
		if err := lorawan.SendUplink(payload, session); err != nil {
			println("error sending data:", err)
		} else {
			print("correctly sent: ")
			for _, b := range payload {
				print(b, " ")
			}
			println()
		}

		println("done! waiting for the next one...")

		time.Sleep(LORAWAN_UPLINK_DELAY)

		onboardLed.Low()

		time.Sleep(1 * time.Second)
	}
}

func connectLoRaWAN(otaa *lorawan.Otaa, session *lorawan.Session) error {
	start := time.Now()
	var err error
	for time.Since(start) < LORAWAN_JOIN_TIMEOUT {
		println("Trying to join network")
		err = lorawan.Join(otaa, session)
		if err == nil {
			println("Connected to network !")
			return nil
		}
		println("Join error:", err, "retrying in", LORAWAN_RECONNECT_DELAY, "sec")
		time.Sleep(LORAWAN_RECONNECT_DELAY)
	}

	err = errors.New("Unable to join Lorawan network")
	println(err.Error())
	return err
}
