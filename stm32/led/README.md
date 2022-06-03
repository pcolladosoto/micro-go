# Driving the on-board LED
This program drives the on-board LED on STM32 boards.

This LED is connected to the board's digital `pin C13` and it's high-level-active. That is, it'll be on when `pin C13` is *HIGH* and viceversa.

This project is practically the same as the one provide don [TinyGo's Blinky Tutorial](https://tinygo.org/docs/tutorials/blinky/): we just added some comments to further clarify things.

It's main use is making sure the environment is correctly set up before tackling some more complex projects.

## Wiring
This project requires no additional components to work: we're just driving the on-board LED!

## Compiling and flashing
We don't really need to do anything special. We can just run:

    collado@hoth:0:~/micro-go/stm32/led$ tinygo flash -target=bluepill
    Open On-Chip Debugger 0.11.0
    Licensed under GNU GPL v2
    For bug reports, read
        http://openocd.org/doc/doxygen/bugs.html
    WARNING: interface/stlink-v2.cfg is deprecated, please switch to interface/stlink.cfg
    Info : auto-selecting first available session transport "hla_swd". To override use 'transport select <transport>'.
    Info : The selected transport took over low-level target control. The results might differ compared to plain JTAG/SWD
    Info : clock speed 1000 kHz
    Info : STLINK V2J17S4 (API v2) VID:PID 0483:3748
    Info : Target voltage: 3.196321
    Info : stm32f1x.cpu: hardware has 6 breakpoints, 4 watchpoints
    Info : starting gdb server for stm32f1x.cpu on 3333
    Info : Listening on port 3333 for gdb connections
    target halted due to debug-request, current mode: Thread 
    xPSR: 0x01000000 pc: 0x08000e8c msp: 0x20000800
    ** Programming Started **
    Info : device id = 0x20036410
    Info : flash size = 64kbytes
    ** Programming Finished **
    ** Resetting Target **
    shutdown command invoked
