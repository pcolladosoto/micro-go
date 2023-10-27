# Driving the on-board LED
This program drives the on-board LED on Arduino boards.

This LED is connected to the board's digital `pin 13` and it's high-level-active. That is, it'll be on when `pin 13` is *HIGH* and viceversa.

This project is practically the same as the one provide don [TinyGo's Blinky Tutorial](https://tinygo.org/docs/tutorials/blinky/): we just added some comments to further clarify things.

It's main use is making sure the environment is correctly set up before tackling some more complex projects.

## Wiring
This project requires no additional components to work: we're just driving the on-board LED!

## Compiling and flashing
We don't really need to do anything special. We can just run:

    collado@hoth:0:~/go-micro/arduino/led$ tinygo flash -target=arduino

    avrdude: AVR device initialized and ready to accept instructions

    Reading | ################################################## | 100% 0.00s

    avrdude: Device signature = 0x1e950f (probably m328p)
    avrdude: NOTE: "flash" memory has been specified, an erase cycle will be performed
            To disable this feature, specify the -D option.
    avrdude: erasing chip
    avrdude: reading input file "/var/folders/d3/z19cz5f50xzd310hg9ygp7gc0000gn/T/tinygo987352048/main.hex"
    avrdude: writing flash (2460 bytes):

    Writing | ################################################## | 100% 0.41s

    avrdude: 2460 bytes of flash written
    avrdude: verifying flash memory against /var/folders/d3/z19cz5f50xzd310hg9ygp7gc0000gn/T/tinygo987352048/main.hex:

    Reading | ################################################## | 100% 0.33s

    avrdude: 2460 bytes of flash verified

    avrdude done.  Thank you.
