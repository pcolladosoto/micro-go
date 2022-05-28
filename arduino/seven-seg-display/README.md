# Driving a 7-segment display
This program drives a [7-segment display](https://en.wikipedia.org/wiki/Seven-segment_display) connected to the digital pins of an Arduino board.

## Our display
The display we have used is [*KingBright's SA56-11EWA*](https://uk.farnell.com/kingbright/sa56-11ewa/display-seven-segment-19-05mm/dp/1142440) ([datasheet](https://www.farnell.com/datasheets/1864136.pdf)). This display is [**low-level-active**](https://en.wikipedia.org/wiki/Logic_level#Active_state), that is, the common connection should be connected to a high level pin and segments are then activated with a low level.

Should your display be *high-level-active* you should remove the logical `NOT` operator (i.e. `!`) on `line 57` of `main.go`.

## Wiring
The following table contains the necessary wiring. Note we're using the printed pin names on an Arduino Uno board:

    + --------------------- +
    | Display | Arduino Uno |
    + --------|------------ +
    |    1    |  Digital 2  |
    |    2    |  Digital 5  |
    |    3    |  Power 3.3V |
    |    4    |  Digital 4  |
    |    5    |    None     |
    |    6    |  Digital 3  |
    |    7    |  Digital 2  |
    |    8    |    None     |
    |    9    |  Digital 7  |
    |   10    |  Digital 8  |
    + --------------------- +

## Compiling and flashing
We don't really need to do anything special. We can just run:

    collado@hoth:0:~/go-micro/arduino/seven-seg-display$ tinygo flash -target=arduino 

    avrdude: AVR device initialized and ready to accept instructions

    Reading | ################################################## | 100% 0.00s

    avrdude: Device signature = 0x1e950f (probably m328p)
    avrdude: NOTE: "flash" memory has been specified, an erase cycle will be performed
            To disable this feature, specify the -D option.
    avrdude: erasing chip
    avrdude: reading input file "/var/folders/d3/z19cz5f50xzd310hg9ygp7gc0000gn/T/tinygo2289048706/main.hex"
    avrdude: writing flash (3348 bytes):

    Writing | ################################################## | 100% 0.56s

    avrdude: 3348 bytes of flash written
    avrdude: verifying flash memory against /var/folders/d3/z19cz5f50xzd310hg9ygp7gc0000gn/T/tinygo2289048706/main.hex:

    Reading | ################################################## | 100% 0.44s

    avrdude: 3348 bytes of flash verified

    avrdude done.  Thank you.
