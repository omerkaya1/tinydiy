package buzzer

import (
	"machine"
	"time"
)

type Buzzer struct {
	pin machine.Pin
}

func New(pin machine.Pin) *Buzzer {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &Buzzer{pin: pin}
}

type BeepParams struct {
	Count    int
	Duration time.Duration
	Delay    time.Duration
}

func (b *Buzzer) Beep(p BeepParams) {
	for range p.Count {
		b.pin.High()
		time.Sleep(p.Duration)
		b.pin.Low()
		if p.Delay > 0 {
			time.Sleep(p.Delay)
		}
	}
}
