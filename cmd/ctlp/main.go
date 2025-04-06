// ctlp - trafic lights + pedestrian crossing symulator
// Used components:
// Traffic part
// - 3 leds - (red, green, blue - GPIO - D13, D12, D11)
// - Resistors (220 ohm * 3)
// Pedestrian part
// - 2 leds - (red, green - GPIO - D4, D5)
// - Resistors (220 ohm * 2)
// - button (GPIO - D2)
package main

import (
	"machine"
	"sync"
	"time"
)

type crossingState struct {
	mx    sync.RWMutex
	state bool
}

func (s *crossingState) getState() bool {
	s.mx.Lock()
	state := s.state
	s.mx.Unlock()
	return state
}

func (s *crossingState) setState(state bool) {
	s.mx.Lock()
	s.state = state
	s.mx.Unlock()
}

func pedestrians(btn *machine.Pin, state *crossingState) {
CHECK_BUTTON_STATE:
	if btn.Get() {
		state.setState(true)
	}
	time.Sleep(time.Millisecond * 10)
	goto CHECK_BUTTON_STATE
}

func startCrossing(state *crossingState, red, green *machine.Pin) {
	red.Low()
	green.High()
	time.Sleep(time.Millisecond * 1000)

	state.setState(false)

	red.High()
	green.Low()
}

func main() {
	// state
	var state crossingState

	// pedestrian part
	outCfg := machine.PinConfig{Mode: machine.PinOutput}
	pGreen := machine.D4
	pRed := machine.D5

	pGreen.Configure(outCfg)
	pRed.Configure(outCfg)

	btn := machine.D2
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	go pedestrians(&btn, &state)

	// traffic part
	red := machine.D13
	yellow := machine.D12
	green := machine.D11

	red.Configure(outCfg)
	yellow.Configure(outCfg)
	green.Configure(outCfg)

	// initial state
	pRed.High()

RED:
	red.High()
	green.Low()
	for range 2900 {
		if state.getState() {
			startCrossing(&state, &pRed, &pGreen)
			break
		}
		time.Sleep(time.Millisecond)
	}

YELLOW:
	yellow.High()
	time.Sleep(time.Millisecond * 200)

	red.Low()
	yellow.Low()
	green.High()

	for range 1200 {
		if state.getState() {
			red.Low()
			green.Low()
			yellow.High()
			time.Sleep(time.Millisecond * 200)
			red.High()
			green.Low()
			yellow.Low()

			startCrossing(&state, &pRed, &pGreen)
			goto YELLOW
		}
		time.Sleep(time.Millisecond)
	}

	green.Low()
	yellow.High()
	time.Sleep(time.Millisecond * 200)

	yellow.Low()
	goto RED
}
