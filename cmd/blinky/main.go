package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	cfg := machine.PinConfig{Mode: machine.PinOutput}

	externalLed := machine.D13
	externalLed.Configure(cfg)

	for {
		led.High()
		externalLed.High()

		time.Sleep(time.Millisecond * 1000)

		externalLed.Low()
		led.Low()

		time.Sleep(time.Millisecond * 200)

		led.High()
		externalLed.High()

		time.Sleep(time.Millisecond * 500)

		led.Low()
		externalLed.Low()
	}
}
