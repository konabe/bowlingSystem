package pinNumbers

//PinNumbers ピンの番号を指定する場合に使用します。不正なデータの場合はnilが返却されます。
type PinNumbers struct {
	Values []int
}

func New(numbers ...int) PinNumbers {
	return PinNumbers{
		Values: makeUniqueAndValidNumbers(numbers),
	}
}

func (pinNumbers PinNumbers) Len() int {
	return len(pinNumbers.Values)
}

func makeUniqueAndValidNumbers(numbers []int) []int {
	bucket := [10]bool{}
	newNumbers := []int{}
	for _, v := range numbers {
		if 1 <= v && v <= 10 {
			bucket[v-1] = true
		}
	}
	for i, v := range bucket {
		if v {
			newNumbers = append(newNumbers, i+1)
		}
	}
	return newNumbers
}
