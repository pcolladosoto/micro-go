# Running Go on MicroControllers
The world of micro-controllers (i.e. Arduino and co.) has long been dominated by languages such as C and vendor-specific options.

Through working with it, we've come to fall in love with Go, so we decided to try and make it work on embedded platforms with strict requirements on what can be run on them.

It was during this search that we came across the [TinyGo](https://tinygo.org) project. This repository contains several code projects targeting some popular micro-controllers which are always compiled with `tinygo`.

## Getting ready
In order to use the `tinygo` compiler we need to set some things up first. We are using macOS, so we'll include installation instructions for it as a goodie. However, you should always abide by what's written on [TinyGo's Installation Instructions](https://tinygo.org/getting-started/install/) should that differ from the following. You can also find the installation instructions for other platforms there!

Please bear in mind that you should have installed [Go v1.15+](https://go.dev) (note version `1.17+` is preferred) before trying to use `tinygo`. Go's installation is beyond the scope of this document, but you'll find information on how to install it no matter your platform [here](https://go.dev/doc/install).

### Installing on macOS
We assume you're using [Homebrew](https://brew.sh) to install packages. If not... tough luck! You can try to look into how to [compile TinyGo from scratch](https://tinygo.org/docs/guides/build/) though.

We should begin by adding a new *tap* to Homebrew. A tap is just a directory for new goodies as seen on [`brew(1)`](https://docs.brew.sh/Manpage). After adding the new origin where software can be obtained from we can just install TinyGo and we're done!

    # Let's add a TinyGo's tap
    brew tap tinygo-org/tools

    # And install TinyGo right away!
    brew install tinygo

We can now check TinyGo is installed by running `tinygo version`:

    collado@hoth:0:~$ tinygo version
    tinygo version 0.23.0 darwin/amd64 (using go version go1.17.3 and LLVM version 14.0.0)

That's it... unless we are dealing with some specific micro-controllers. A prime example requiring a bit more of work are AVR-based Arduino boards such as the Arduino Uno. Not to worry though: we'll cover needed dependencies in `README.md`s found within each subdirectory in the repository.

## What now?
At this point you should be capable of leveraging `tinygo` to compile and flash any of the [supported boards](https://tinygo.org/docs/reference/microcontrollers/). We have divided our examples in a series of directories where each targets a specific hardware platform. Feel free to browse around!

## Working with TinyGo
When working with TinyGo we'll mostly be interacting with the `machine` package it provides. This package implements the 'link' between software and hardware. You can find a general discussion on its components [here](https://tinygo.org/docs/reference/machine/), but rest assured we'll always provide a link to the specifics of the package for each of the board types we work on.

TinyGo is always invoked through the `tinygo` command. We'll provide the commands you'll commonly be using, but you can find a full-fledged reference [here](https://tinygo.org/docs/reference/usage/).

## Writing code on Visual Studio Code
We write our code on [Visual Studio Code](https://code.visualstudio.com). Given the `machine` package is not part of Go's standard libraries we looked around for extensions offering autocomplete for said package. Luckily for us, we quickly found the [TinyGo extension](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo).

What it basically does is alter your Go environment (i.e. `GOROOT` and `GOFLAGS`) so that other extensions can find the `machine` package implementation. That implies this extension relies on the stock [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go).

These extensions are by no means necessary, but we do find them extremely useful.

## Contributing
This repository's contents are rather scarce: we would love to include new contributions! What's more, we know we're not perfect, so feel free to open an issue if something doesn't work just right: we'll be happy to help.
