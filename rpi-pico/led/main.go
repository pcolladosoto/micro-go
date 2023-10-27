package main

/* Please note this example is taken from TinyGo's tutorials:
 *     https://tinygo.org/docs/tutorials/blinky/
 */

import (
	"machine"
	"time"
	"math/rand"
)

func main() {
	// Let's get a reference to the on-board LED
	onboardLed := machine.LED

	// We need to initialise it as an output
	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// This for loop will run indefinitely!
	for {
		// We'll turn the on-board LED off...
		onboardLed.Low()

		// ... and wait 500 ms
		time.Sleep(time.Millisecond * time.Duration(500 + rand.Intn(500)))

		// Turn it back on...
		onboardLed.High()

		// ... and wait another 500 ms
		time.Sleep(time.Millisecond * time.Duration(500 + rand.Intn(500)))
	}
}
