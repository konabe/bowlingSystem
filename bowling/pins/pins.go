package pins

import (
	pinNumbers "bowlingSystem/bowling/pins/numbers"
	"errors"
)

type Pins struct {
	values [10]bool
}

func New() *Pins {
	return &Pins{values: [10]bool{true, true, true, true, true, true, true, true, true, true}}
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

func (pins Pins) GetRestNumbers() []int {
	var resultPins []int
	for i, v := range pins.values {
		if v {
			resultPins = append(resultPins, i+1)
		}
	}
	return resultPins
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
	if pn == nil {
		return errors.New("numbersが有効ではありません")
	}
	for _, v := range pn.Values {
		pins.values[v-1] = false
	}
	return nil
}
