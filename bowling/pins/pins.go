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

func (pins Pins) KnockDown(numbers []int) (*Pins, error) {
	newPins := copy(pins)
	err := newPins.knockDown(numbers)
	if err == nil {
		return newPins, nil
	} else {
		return nil, err
	}
}

func (pins Pins) GetRestNumbers() pinNumbers.PinNumbers {
	var resultPins []int
	for i, v := range pins.values {
		if v {
			resultPins = append(resultPins, i+1)
		}
	}
	return pinNumbers.New(resultPins...)
}

func copy(pins Pins) *Pins {
	newPins := New()
	for i, v := range pins.values {
		newPins.values[i] = v
	}
	return newPins
}

func (pins *Pins) knockDown(numbers []int) error {
	pn := pinNumbers.New(numbers...)
	for _, v := range pn.Values {
		pins.values[v-1] = false
	}
	return nil
}
