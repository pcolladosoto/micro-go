# Implementing a LoRaWAN stack on a Raspberry Pi Pico
The good news is most of the work is done: [`tinygo-org/drivers`](https://github.com/tinygo-org/drivers/tree/release/examples/lora/lorawan/basic-demo)
provides a full-fledged LoRaWAN stack implemented in Go and, what's better, offers examples on how to to leverage it.

This directory contains a tentaive implementation for our use case: we will try to integrate it with a Dragino LoRaWAN
gateway to then receive data on a TTN application server.

Let's get to work!
