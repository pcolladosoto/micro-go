package main

import (
	"time"

	"machine"
)

const (
	START_DELAY int = 10

	UART_BAUDRATE uint32             = 9600
	UART_DATABITS uint8              = 8
	UART_STOPBITS uint8              = 1
	UART_PARITY   machine.UARTParity = machine.ParityNone
)

func main() {
	// machine.InitSerial()

	for i := 0; i < START_DELAY; i++ {
		println("starting in", START_DELAY-i, "seconds...")
		time.Sleep(1 * time.Second)
	}

	uart := machine.UART1

	// Configure already sets the 8/1/N options as seen on:
	// https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040_uart.go
	if err := uart.Configure(machine.UARTConfig{
		BaudRate: UART_BAUDRATE,
		TX:       machine.GP4,
		RX:       machine.GP5,
	}); err != nil {
		println("error configuring the UART:", err)
	}

	// Channel 0 voltage on N4AIA04
	modbusMsg := []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x01}

	var crc crc
	crc.init()
	crc.add(modbusMsg)

	modbusMsg = append(modbusMsg, crc.value()...)
	print("we'll write: ")
	for _, b := range modbusMsg {
		print(b, " ")
	}
	println()

	recvBuff := make([]byte, 10)

	onboardLed := machine.LED

	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	if uart.Buffered() > 0 {
		discard, _ := uart.ReadByte()
		println("discarded initial stuck byte:", discard)
	}

	var dummyCnt uint16 = 0
	for {
		onboardLed.High()

		println("writing data to the UART for the", dummyCnt, "th time...")
		n, err := uart.Write(modbusMsg)
		if err != nil {
			println("error writing to the UART:", err)
			continue
		}
		print("wrote ", n, " bytes to the UART: ")
		for _, b := range modbusMsg {
			print(b, " ")
		}
		println()

		nDelays := 0
		for nDelays < 5 {
			if uart.Buffered() > 0 {
				n, err = uart.Read(recvBuff)
				if err != nil {
					println("error receiving data form the UART:", err)
					continue
				}
				print("received ", n, " bytes: ")
				for _, b := range recvBuff[:n] {
					print(b, " ")
				}
				println()
				break
			}
			time.Sleep(1 * time.Second)
			nDelays++
		}

		if nDelays == 5 {
			println("timeout on reception...")
		}

		dummyCnt++

		time.Sleep(time.Second * 5)

		onboardLed.Low()

		time.Sleep(1 * time.Second)
	}
}
