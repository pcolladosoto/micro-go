package main

import (
	"machine"
	"time"
)

var (
	/* This table contains the value we should set on each pin driving the 7-segment display.
	 * Note `true` equates to a logical HIGH and `false` produces a logical LOW. The number
	 * displayed by a combination is given by it's index within the containing array. That is,
	 * the first array displays a 0, the next one a 1 and so on.
	 */
	n_to_seg [10][7]bool = [10][7]bool{
		{true, true, true, true, true, true, false},
		{false, true, true, false, false, false, false},
		{true, true, false, true, true, false, true},
		{true, true, true, true, false, false, true},
		{false, true, true, false, false, true, true},
		{true, false, true, true, false, true, true},
		{true, false, true, true, true, true, true},
		{true, true, true, false, false, false, false},
		{true, true, true, true, true, true, true},
		{true, true, true, true, false, true, true},
	}
)

func main() {
	/* We need to get the references to the pins we're to use:
	 *     onboardLed: Arduino's on-board LED. We blink it N times before displaying number N on the display.
	 *     segs:       This array (it's not a slice!) allows us to control the pins driving each of the display's segments.
	 */
	onboardLed := machine.LED
	segs := [7]machine.Pin{machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7, machine.D8}

	// We need to initialise all of them as outputs: we'll be writing to them.
	onboardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for _, pin := range segs {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// As we omit the condition check, this loop will run indefinitely
	for i := 0; ; i++ {

		// We'll begin by blinking the on-board LED to know what number should be displayed next.
		for j := 0; j < i%10; j++ {
			onboardLed.High()
			time.Sleep(time.Millisecond * 125)
			onboardLed.Low()
			time.Sleep(time.Millisecond * 125)
		}

		/* Then, we just need to write the appropriate value to each pin controlling the display.
		 * Note the logical NOT (i.e. `!`) on line 57: it's due to the display being low-level-active!
		 */
		for seg, v := range n_to_seg[i%10] {
			segs[seg].Set(!v)
		}

		// And finally wait a second before displaying the next number
		time.Sleep(time.Second)
	}
}
