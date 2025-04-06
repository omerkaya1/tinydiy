// Prerequisites:
// - install PuTTy: brew install PuTTy;
// -
package main

import (
	"machine"

	"github.com/omerkaya1/tinydiy/internal/driver/keyboard"
)

func main() {
	keyboard := keyboard.New(
		machine.D2, machine.D3, machine.D4, machine.D5,
		machine.D6, machine.D7, machine.D8, machine.D9,
	)
	for {
		k := keyboard.Key()
		if k != "" {
			println(k)
		}
	}
}
