// beep - beep on button press
// Used components:
// - D2 button
// - D3 beepper
package main

import (
	"machine"
	"time"
)

func main() {
	// init output config
	outCfg := machine.PinConfig{Mode: machine.PinOutput}

	// periferal
	beepper := machine.D3
	beepper.Configure(outCfg)

	inCfg := machine.PinConfig{Mode: machine.PinInput}

	btn := machine.D2
	btn.Configure(inCfg)

	for {
		beepper.Low()
		if btn.Get() {
			beepper.High()
			time.Sleep(time.Millisecond * 5000)
			beepper.Low()
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
