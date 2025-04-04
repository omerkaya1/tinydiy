package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(time.Millisecond * 1000)
		led.Low()
		time.Sleep(time.Millisecond * 50)
		led.High()
		time.Sleep(time.Millisecond * 50)
		led.Low()
	}
}
