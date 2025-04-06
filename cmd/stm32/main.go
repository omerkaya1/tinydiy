//go:build cortexm || baremetal || linux || arm || nucleof722ze || stm32f7x2 || stm32f7 || stm32 || tinygo || purego || osusergo || math_big_pure_go || gc.conservative || scheduler.tasks || serial.uart
// +build cortexm baremetal linux arm nucleof722ze stm32f7x2 stm32f7 stm32 tinygo purego osusergo math_big_pure_go gc.conservative scheduler.tasks serial.uart

package main

import (
	"machine"
	"time"
)

func main() {
	// init output config
	outCfg := machine.PinConfig{Mode: machine.PinOutput}

	// periferal
	blue := machine.LED_BLUE
	red := machine.LED_RED
	green := machine.LED_GREEN

	blue.Configure(outCfg)
	red.Configure(outCfg)
	green.Configure(outCfg)

	for {
		blue.High()
		time.Sleep(time.Second * 1)
		blue.Low()
		time.Sleep(time.Second * 1)
		red.High()
		time.Sleep(time.Second * 1)
		red.Low()
		time.Sleep(time.Second * 1)
		green.High()
		time.Sleep(time.Second * 1)
		green.Low()
		time.Sleep(time.Second * 1)
	}
}
