# STM32 Boards
This directory contains several examples of programs capable of running on STM32 boards.

Please note we've **only tested** them on an a [**bluepill**](https://stm32duinoforum.com/forum/wiki_subdomain/index_title_Blue_Pill.html): it's the only board we have lying around. To the extent of what we know these should work on other STM32-based boards, but we cannot be sure.

As hinted in the root [`README.md`](../README.md), we need a few extra dependencies to work with [STM32](https://en.wikipedia.org/wiki/STM32)-based SoCs. The main difference with respect to Arduino boards is that we need to acquire a *programmer*. This is just a dongle offering a USB-A (i.e. the regular USB) connection on one end and exposing a set of pins on the other. Instead of connecting the board directly to our machine's USB port we'll connect it to the programmer and then connect the programmer to our machine. The connection involves 4-pins: don't worry about it too much!

We'll provide the installation instructions for software dependencies for macOS. You can refer to [OpenOCD](https://openocd.org/pages/getting-openocd.html)'s site for information on other platforms.

## Getting STM32 dependencies on macOS
As stated above, we'll need to grab some software for flashing the board before moving on to the programmer.

You can refer to [this site](https://alexbirkett.github.io/microcontroller/2019/03/30/flash_bluepill_using_ST_link.html) for a more in-depth explanation of the process below and some goodies on communicating with the board through [`telnet(1)`](https://linux.die.net/man/1/telnet).

### The flashing software
Flashing is done with [`openocd`](https://openocd.org). We just need to install it and that's that! The good news is it's available on Homebrew's default taps:

    # We just need to install `openocd`!
    brew install openocd

Just like with Arduino, TinyGo will 'hide' `openocd` from us: we won't need to call it directly!

### Wiring the programmer up
The programmer we're using is the [*ST-Link v2*](https://stm32-base.org/boards/Debugger-STM32F103C8U6-STLINKV2). It leverages ARM's [Serial Wire Debug](https://wiki.segger.com/index.php?title=SWD) interface to communicate with the board, so we just need to work with 4 pins. Bear in mind the following wiring table **applies to our own model**: it might be different than your version! Always check the documentation provided by your buyer.

The connections you have to make are:

    + --------------------- +
    | Programmer | BluePill |
    + -----------|--------- +
    |  2 - SWCLK |  SWCLK   |
    |  4 - SWDIO |  SWIO    |
    |  6 -   GND |  3V3     |
    |  8 -  3.3V |  GND     |
    + --------------------- +

The pins on the BluePill are not on the rails: the're exposed on the opposite side of the Micro-USB socket. You can also take a look at [this video](https://www.youtube.com/watch?v=KgR3uM21y7o) for a more visual explanation.

Once you connect the programmer to a USB socket you should see the power LED (i.e. the red one) light up!

If you encounter errors when flashing the board regarding connectivity make sure to double check the wiring between the board and programmer!

## Interacting with the board
Just like we said on the root [`README.md`](../README.md), we'll interact with the boards themselves though the `machine` package. In other words, we can 'only' do as much as the `machine` package allows us to. That's why it's crucial to know what we can and cannot do.

This is answered in the form of a bunch of documentation you can query [here](https://tinygo.org/docs/reference/microcontrollers/machine/bluepill/). However, if you've enabled autocompletion though the TinyGo VS Code extension you can 'read' all that information interactively through the autocomplete prompts. That's one of the main reasons we find it so useful!

In any case, it's a good idea to familiarise yourself with the functions and types you'll be able to use to interact with Arduino boards.

## Creating a new project
We're using Go [modules](https://go.dev/ref/mod). That means each project is contained inside its own directory. Thus, to start a new project we need to begin by creating a new directory:

    # Let's create the led directory...
    mkdir led

    # ... and navigate to it
    cd led

Next, we need to initialise a new module. We can call it whatever we want (we've chosen `stm32-led`):

    go mod init stm32-led

Finally, we can begin writing our code on a regular file. We follow the convention that the `main()` function is always defined on a file called `main.go`. That's where everything starts! Please note the code **needs** to belong to the `main` package (i.e. the initial line should be `package main`).

## Compiling stuff
Compilation is done through `tinygo`. We can compile the source code with a `tinygo build`, but we'll usually want to compile and upload the code to our board in a single step:

    # This will compile our code for Arduino boards and then upload it right away!
    tinygo flash -target=bluepill

Please note that for this to work you **need** to be on the **source code directory**.
