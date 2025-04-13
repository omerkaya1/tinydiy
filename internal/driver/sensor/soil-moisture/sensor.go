package soilmoisture

import "machine"

const (
	HW080ArduinoDry = 65472
	HW080ArduinoWet = 14500
)

const totalCategories = 6

type MoistureLevel uint8

const (
	CompletelyDry = iota
	Dry
	SlightlyDry
	Moist
	VeryMoist
	Water
)

type Sensor struct {
	category     uint16
	dryThreshold uint16
	wetThreshold uint16
	adc          *machine.ADC
	voltage      machine.Pin
}

type Params struct {
	DryThreshold uint16
	WetThreshold uint16
	Voltage      machine.Pin
}

func New(params Params) *Sensor {
	// derive categories
	category := (params.DryThreshold - params.WetThreshold) / totalCategories
	s := Sensor{
		category:     category,
		dryThreshold: params.DryThreshold,
		wetThreshold: params.WetThreshold,
		voltage:      params.Voltage,
		adc:          &machine.ADC{Pin: params.Voltage},
	}
	s.voltage.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.adc.Configure(machine.ADCConfig{})
	return &s
}

func (s *Sensor) On() {
	s.voltage.High()
}

func (s *Sensor) Off() {
	s.voltage.Low()
}

func (s *Sensor) Read() MoistureLevel {
	val := s.adc.Get()
	switch {
	case val >= s.dryThreshold:
		return CompletelyDry
	case val >= s.dryThreshold-s.category:
		return Dry
	case val >= s.dryThreshold-s.category*2:
		return SlightlyDry
	case val >= s.dryThreshold-s.category*3:
		return Moist
	case val >= s.dryThreshold-s.category*4:
		return VeryMoist
	default:
		return Water
	}
}
