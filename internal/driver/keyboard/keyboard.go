package keyboard

import "machine"

type Keyboard struct {
	enabled bool

	lastCol int
	lastRow int

	cols [4]machine.Pin
	rows [4]machine.Pin

	mapping [4][4]string
}

func New(
	r1, r2, r3, r4,
	c1, c2, c3, c4 machine.Pin,
) *Keyboard {
	inCfg := machine.PinConfig{Mode: machine.PinInput}
	outCfg := machine.PinConfig{Mode: machine.PinOutput}

	r1.Configure(outCfg)
	r2.Configure(outCfg)
	r3.Configure(outCfg)
	r4.Configure(outCfg)

	c1.Configure(inCfg)
	c2.Configure(inCfg)
	c3.Configure(inCfg)
	c4.Configure(inCfg)

	return &Keyboard{
		enabled: true,

		lastCol: -1,
		lastRow: -1,

		cols: [4]machine.Pin{
			c1, c2, c3, c4,
		},
		rows: [4]machine.Pin{
			r1, r2, r3, r4,
		},

		mapping: [4][4]string{
			{"1", "2", "3", "A"},
			{"4", "5", "6", "B"},
			{"7", "8", "9", "C"},
			{"*", "0", "#", "D"},
		},
	}
}

func (k *Keyboard) Key() string {
	row, col := k.Read()
	if row == -1 || col == -1 {
		return ""
	}
	return k.mapping[row][col]
}

func (k *Keyboard) Read() (int, int) {
	for i := range k.rows {
		row := k.rows[i]
		row.Low()

		for j := range k.cols {
			col := k.cols[j]

			if !col.Get() && k.enabled {
				k.enabled = false

				k.lastCol = j
				k.lastRow = i

				row.High()
				return k.lastRow, k.lastCol
			}

			if col.Get() &&
				j == k.lastCol &&
				i == k.lastRow &&
				!k.enabled {
				k.enabled = true
			}
		}

		row.High()
	}

	return -1, -1
}
