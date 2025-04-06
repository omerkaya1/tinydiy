// ctl - trafic lights symulator
// Used components:
// - 3 leds - (red, green, blue - GPIO - D13, D12, D11)
// - Resistors (220 ohm * 3)
package main

import (
	"machine"
	"sync"
	"time"
)

const (
	greenDuration  = 1200 * time.Millisecond
	yellowDuration = 200 * time.Millisecond
	redDuration    = 2900 * time.Millisecond
)

type crossingState struct {
	mx    sync.RWMutex
	state bool
}

const (
	crossingNotStarted = iota
	crossingRequested
	crossingInProgress
	crossingFinished
)

type streetlight struct {
	red    machine.Pin
	yellow machine.Pin
	green  machine.Pin
}

// func new() streetlight {
// 	return streetlight{
// 		red:    machine.D13,
// 		yellow: machine.D12,
// 		green:  machine.D11,
// 	}
// }

// func (sl streetlight) cfg() {
// 	outCfg := machine.PinConfig{Mode: machine.PinOutput}

// 	sl.red.Configure(outCfg)
// 	sl.yellow.Configure(outCfg)
// 	sl.green.Configure(outCfg)
// }

// func (sl streetlight) run() {
// 	sl.red.High()
// 	time.Sleep(time.Millisecond * 2900)
// 	sl.red.Low()
// 	sl.yellow.High()
// 	time.Sleep(time.Millisecond * 200)
// 	sl.green.High()
// 	sl.yellow.Low()
// 	time.Sleep(time.Millisecond * 1200)
// 	sl.green.Low()
// 	sl.yellow.High()
// 	time.Sleep(time.Millisecond * 200)
// 	sl.yellow.Low()
// }

func traffic(state *crossingState) {
	// outCfg := machine.PinConfig{Mode: machine.PinOutput}
	// red := machine.D13
	// yellow := machine.D12
	// green := machine.D11

	// red.Configure(outCfg)
	// yellow.Configure(outCfg)
	// green.Configure(outCfg)

	// for {
	// 	// we check the state each time we start
	// 	red.High()
	// 	// we either sleep here for the total of the specific duration, or we deviate for a pedestrian press and then come back

	// 	state.mx.RLock()
	// 	data := state.state == crossingRequested
	// 	state.mx.RUnlock()
	// 	if data {
	// 		state.state = crossingInProgress

	// 		// pedestrians(state)
	// 		state.mx.Lock()
	// 		state.state = crossingFinished
	// 		state.mx.Unlock()
	// 	} else {
	// 		state.mx.RUnlock()
	// 	}

	// 	time.Sleep(redDuration)

	// 	red.Low()
	// 	yellow.High()
	// 	time.Sleep(yellowDuration)

	// 	green.High()
	// 	yellow.Low()
	// 	time.Sleep(greenDuration)

	// 	green.Low()
	// 	yellow.High()
	// 	time.Sleep(yellowDuration)

	// 	yellow.Low()
	// }
}

func pedestrians(btn *machine.Pin, state *crossingState) {
HERE:
	if btn.Get() {
		state.mx.Lock()
		state.state = true
		state.mx.Unlock()
	}
	time.Sleep(time.Millisecond * 10)
	goto HERE
}

func startCrossing(state *crossingState, red, green *machine.Pin) {
	red.Low()
	green.High()
	time.Sleep(time.Millisecond * 1000)

	state.mx.Lock()
	state.state = false
	state.mx.Unlock()

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
		state.mx.RLock()
		data := state.state
		state.mx.RUnlock()

		if data {
			// red.Low()
			// green.Low()
			// yellow.High()
			// time.Sleep(time.Millisecond * 2000)
			// red.High()
			// green.Low()
			// yellow.Low()

			startCrossing(&state, &pRed, &pGreen)
			goto YELLOW
		}
		time.Sleep(time.Millisecond)
	}

YELLOW:
	yellow.High()
	time.Sleep(time.Millisecond * 200)

	// GREEN:
	red.Low()
	yellow.Low()
	green.High()

	for range 1200 {
		state.mx.RLock()
		data := state.state
		state.mx.RUnlock()

		if data {
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
