package pump

import (
	"machine"
	"time"
)

type Pump struct {
	pin machine.Pin
}

func New(pin machine.Pin) *Pump {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &Pump{pin: pin}
}

type PourParams struct {
	Duration time.Duration
	Delay    time.Duration
	Count    int
}

func (p *Pump) Pour(params PourParams) {
	for range params.Count {
		p.pin.High()
		time.Sleep(params.Duration)
		p.pin.Low()
		if params.Delay > 0 {
			time.Sleep(params.Delay)
		}
	}
}
