// beep - beep on button press
// Used components:
// - D2 button
// - D3 beepper
package main

import (
	"machine"
	"time"

	"github.com/omerkaya1/tinydiy/internal/driver/buzzer"
)

func main() {
	// periferal
	beepper := buzzer.New(machine.D2)
	for {
		beepper.Beep(buzzer.BeepParams{
			Count:    3,
			Duration: time.Millisecond * 1000,
			Delay:    time.Millisecond * 500,
		})
		time.Sleep(time.Millisecond * 5000)
	}
}
