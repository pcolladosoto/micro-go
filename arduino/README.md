# Arduino Boards
This directory contains several examples of programs capable of running on Arduino boards.

Please note we've **only tested** them on an **original Arduino Uno**: it's the only board we have lying around. To the extent of what we know these should work on other Arduino flavours, but we cannot be sure.

As hinted in the root [`README.md`](../README.md), we need a few extra dependencies to work with [AVR](https://en.wikipedia.org/wiki/AVR_microcontrollers)-based processors such as Arduino Uno's `ATMEGA328P-PU`.

Once again, we'll provide the installation instructions for macOS. You can find the instructions for other platforms [here](https://tinygo.org/getting-started/install/).

## Getting AVR dependencies on macOS
The process is fairly similar to TinyGo's installation: we need to add a new *tap* and install needed packages:

    # Let's add the tap for AVR goodies
    brew tap osx-cross/avr

    # We need to install a C-to-AVR compiler: namely `avr-gcc`
        # Check https://ccrma.stanford.edu/~juanig/articles/wiriavrlib/AVR_GCC.html for more info!
    brew install avr-gcc

    # We also need to upload the code we generate! That's where `avrdude` comes in.
        # Check https://www.nongnu.org/avrdude/ for more info!
    brew install avrdude

With that we should be ready to go. One of the best things about TinyGo is that it 'hides' calls to necessary tools. That is, you'll not be calling `avr-gcc` and `avrdude` directly. That takes a lot of the complexity of working with micro-controllers out!

## Interacting with the board
Just like we said on the root [`README.md`](../README.md), we'll interact with the boards themselves though the `machine` package. In other words, we can 'only' do as much as the `machine` package allows us to. That's why it's crucial to know what we can and cannot do.

This is answered in the form of a bunch of documentation you can query [here](https://tinygo.org/docs/reference/microcontrollers/machine/arduino/). However, if you've enabled autocompletion though the TinyGo VS Code extension you can 'read' all that information interactively through the autocomplete prompts. That's one of the main reasons we find it so useful!

In any case, it's a good idea to familiarise yourself with the functions and types you'll be able to use to interact with Arduino boards.

## Creating a new project
We're using Go [modules](https://go.dev/ref/mod). That means each project is contained inside its own directory. Thus, to start a new project we need to begin by creating a new directory:

    # Let's create the led directory...
    mkdir led

    # ... and navigate to it
    cd led

Next, we need to initialise a new module. We can call it whatever we want (we've chosen `arduino-led`):

    go mod init arduino-led

Finally, we can begin writing our code on a regular file. We follow the convention that the `main()` function is always defined on a file called `main.go`. That's where everything starts! Please note the code **needs** to belong to the `main` package (i.e. the initial line should be `package main`).

## Compiling stuff
Compilation is done through `tinygo`. We can compile the source code with a `tinygo build`, but we'll usually want to compile and upload the code to our board in a single step:

    # This will compile our code for Arduino boards and then upload it right away!
    tinygo flash -target=arduino

Please note that for this to work you **need** to be on the **source code directory**.
