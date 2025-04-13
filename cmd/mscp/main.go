package main

import (
	"machine"
	"time"

	soilmoisture "github.com/omerkaya1/tinydiy/internal/driver/sensor/soil-moisture"
)

func main() {
	// init ADC
	machine.InitADC()

	// 65472 / 14500
	// use putty to check the serial port output

	sensor := soilmoisture.New(soilmoisture.Params{
		DryThreshold: soilmoisture.HW080ArduinoDry,
		WetThreshold: soilmoisture.HW080ArduinoWet,
		Voltage:      machine.D8,
	})

	for {
		sensor.On()
		val := sensor.Read()

		switch val {
		case soilmoisture.CompletelyDry:
			println("CompletelyDry")
		case soilmoisture.Dry:
			println("Dry")
		case soilmoisture.SlightlyDry:
			println("SlightlyDry")
		case soilmoisture.Moist:
			println("Moist")
		case soilmoisture.VeryMoist:
			println("VeryMoist")
		case soilmoisture.Water:
			println("Water")
		}

		sensor.Off()
		time.Sleep(time.Second * 5)
	}
}
