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

func (pins Pins) KnockDown(numbers []int) (*Pins, error) {
	unique_numbers := util.Unique(numbers)
	if !IsValidNumbers(unique_numbers) {
		return nil, errors.New("numbersが有効ではありません")
	}
	newPins := New()
	for _, v := range unique_numbers {
		newPins.Values[v-1] = false
	}
	return newPins, nil
}

func (pins Pins) GetRestNumbers() []int {
	resultPins := []int{}
	for i, v := range pins.Values {
		if v {
			resultPins = append(resultPins, i+1)
		}
	}
	return resultPins
}

func IsValidNumbers(numbers []int) bool {
	availableNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range numbers {
		if !util.Contains(availableNumbers, v) {
			return false
		}
	}
	return true
}
