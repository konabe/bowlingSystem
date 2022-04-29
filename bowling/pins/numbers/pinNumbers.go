package pinNumbers

import "bowlingSystem/util"

type PinNumbers struct {
	Values []int
}

func New(numbers ...int) *PinNumbers {
	uniqueNumbers := util.Unique(numbers)
	if !isValidNumbers(uniqueNumbers) {
		return nil
	}
	return &PinNumbers{
		Values: uniqueNumbers,
	}
}

func availableNumbers() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func isValidNumbers(numbers []int) bool {
	for _, v := range numbers {
		if !util.Contains(availableNumbers(), v) {
			return false
		}
	}
	return true
}
