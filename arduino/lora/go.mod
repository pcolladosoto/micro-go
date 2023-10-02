module lora

go 1.19

replace github.com/ulbios/lora/sx1276-driver/arduino => ../../../../ulbios/lora/sx1276-driver/arduino

require (
	github.com/ulbios/lora/sx1276-driver/arduino v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.26.0
)

require github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
