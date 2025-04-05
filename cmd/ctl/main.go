// ctl - trafic lights symulator
// Used components:
// - 3 leds - (red, green, blue - GPIO - D13, D12, D11)
// - Resistors (220 ohm * 3)
package main

import (
	"machine"
	"time"
)

type streetlight struct {
	red    machine.Pin
	yellow machine.Pin
	green  machine.Pin
}

func new() streetlight {
	return streetlight{
		red:    machine.D13,
		yellow: machine.D12,
		green:  machine.D11,
	}
}

func (sl streetlight) cfg() {
	cfg := machine.PinConfig{Mode: machine.PinOutput}
	sl.red.Configure(cfg)
	sl.yellow.Configure(cfg)
	sl.green.Configure(cfg)
}

func (sl streetlight) run() {
	sl.red.High()
	time.Sleep(time.Millisecond * 2900)
	sl.red.Low()
	sl.yellow.High()
	time.Sleep(time.Millisecond * 200)
	sl.green.High()
	sl.yellow.Low()
	time.Sleep(time.Millisecond * 1200)
	sl.green.Low()
	sl.yellow.High()
	time.Sleep(time.Millisecond * 200)
	sl.yellow.Low()
}

func main() {
	sl := new()

	sl.cfg()

	for {
		sl.run()
	}
}
