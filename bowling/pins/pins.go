package pins

import (
	"bowlingSystem/util"
	"errors"
)

type Pins struct {
	Values [10]bool
}

func New() *Pins {
	return &Pins{Values: [10]bool{true, true, true, true, true, true, true, true, true, true}}
}

func newWith(pins Pins) *Pins {
	newPins := New()
	for i, v := range pins.Values {
		newPins.Values[i] = v
	}
	return newPins
}

func availableNumbers() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (pins Pins) KnockDown(numbers []int) (*Pins, error) {
	uniqueNumbers := util.Unique(numbers)
	if !isValidNumbers(uniqueNumbers) {
		return nil, errors.New("numbersが有効ではありません")
	}
	newPins := newWith(pins)
	for _, v := range uniqueNumbers {
		newPins.Values[v-1] = false
	}
	return newPins, nil
}

func (pins Pins) GetRestNumbers() []int {
	var resultPins []int
	for i, v := range pins.Values {
		if v {
			resultPins = append(resultPins, i+1)
		}
	}
	return resultPins
}

func isValidNumbers(numbers []int) bool {
	for _, v := range numbers {
		if !util.Contains(availableNumbers(), v) {
			return false
		}
	}
	return true
}
