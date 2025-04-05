// This application configures a microcontroller to turn on an LED when a button is pressed.
// Used components:
// - D13 LED (red)
// - D2 button
// - Resistors (220 ohm and 10k ohm - led and button respectively)
package main

import "machine"

func main() {
	// init output config
	outCfg := machine.PinConfig{Mode: machine.PinOutput}
	inCfg := machine.PinConfig{Mode: machine.PinInput}

	// periferal
	led := machine.D13
	led.Configure(outCfg)

	// buttons
	btn := machine.D2
	btn.Configure(inCfg)

	for {
		if btn.Get() {
			led.High()
		} else {
			led.Low()
		}
	}
}
