package pinNumbers

import "bowlingSystem/util"

//PinNumbers ピンの番号を指定する場合に使用します。不正なデータの場合はnilが返却されます。
type PinNumbers struct {
	Values []int
}

func New(numbers ...int) PinNumbers {
	uniqueNumbers := util.Unique(numbers)
	if !isValidNumbers(uniqueNumbers) {
		return PinNumbers{
			Values: []int{},
		}
	}
	return PinNumbers{
		Values: uniqueNumbers,
	}
}

func (pinNumbers PinNumbers) Len() int {
	return len(pinNumbers.Values)
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
