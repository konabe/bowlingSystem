package pins

import (
	pinNumbers "bowlingSystem/bowling/pins/numbers"
)

type Pins struct {
	values [10]bool
}

func New() *Pins {
	pins := &Pins{}
	for i, _ := range pins.values {
		pins.values[i] = true
	}
	return pins
}

func (pins Pins) KnockDown(numbers []int) *Pins {
	newPins := copy(pins)
	newPins.knockDown(numbers)
	return newPins
}

func (pins Pins) GetRestNumbers() pinNumbers.PinNumbers {
	return pinNumbers.MakeWith(pins.values)
}

func copy(pins Pins) *Pins {
	newPins := New()
	for i, v := range pins.values {
		newPins.values[i] = v
	}
	return newPins
}

func (pins *Pins) knockDown(numbers []int) {
	pn := pinNumbers.New(numbers...)
	for _, v := range pn.Values {
		pins.values[v-1] = false
	}
}
