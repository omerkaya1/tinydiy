package main

import "machine"

func main() {
	// init ADC
	machine.InitADC()

	// 65472 / 14500

	sensor := machine.ADC{Pin: machine.ADC0}
	sensor.Configure(machine.ADCConfig{
		Samples:    4,
		SampleTime: 100,
	})

	for {
		println(sensor.Get())
	}
}
