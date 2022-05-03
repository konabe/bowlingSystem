package pinNumbers

//PinNumbers ピンの番号を指定する場合に使用します。不正なデータの場合はnilが返却されます。
type PinNumbers struct {
	Values []int
}

func New(numbers ...int) PinNumbers {
	return makeUniqueAndValidNumbers(numbers)
}

func MakeWith(bucket [10]bool) PinNumbers {
	newNumbers := []int{}
	for i, v := range bucket {
		if v {
			newNumbers = append(newNumbers, i+1)
		}
	}
	return PinNumbers{
		Values: newNumbers,
	}
}

func (pinNumbers PinNumbers) Len() int {
	return len(pinNumbers.Values)
}

func makeUniqueAndValidNumbers(numbers []int) PinNumbers {
	bucket := [10]bool{}
	for _, v := range numbers {
		if 1 <= v && v <= 10 {
			bucket[v-1] = true
		}
	}
	return MakeWith(bucket)
}
