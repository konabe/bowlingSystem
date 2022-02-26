package frame

import (
	"bowlingSystem/bowling/pins"
	"errors"
)

type TenFrame struct {
	Frame
	ThirdRestPins *pins.Pins
	ThirdScore    int
}

func NewTen() *TenFrame {
	return &TenFrame{
		Frame:         *New(),
		ThirdRestPins: nil,
		ThirdScore:    0,
	}
}

func (frame *TenFrame) BowlSecond(numbers []int) error {
	var currentPins *pins.Pins
	var err error
	if frame.FirstScore == 10 {
		newPins := pins.New()
		currentPins, err = newPins.KnockDown(numbers)
	} else {
		currentPins, err = frame.FirstRestPins.KnockDown(numbers)
	}
	if frame.BowlCount != 1 {
		return errors.New("不正なコールです")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	frame.SecondRestPins = currentPins
	if frame.FirstScore == 10 {
		frame.SecondScore = 10 - len(frame.SecondRestPins.GetRestNumbers())
	} else {
		frame.SecondScore = len(frame.FirstRestPins.GetRestNumbers()) - len(frame.SecondRestPins.GetRestNumbers())
	}
	frame.BowlCount++
	return nil
}

func (frame *TenFrame) BowlThird(numbers []int) error {
	var currentPins *pins.Pins
	var err error
	if frame.SecondScore == 10 || frame.FirstScore+frame.SecondScore == 10 {
		newPins := pins.New()
		currentPins, err = newPins.KnockDown(numbers)
	} else {
		currentPins, err = frame.FirstRestPins.KnockDown(numbers)
	}
	if frame.BowlCount != 2 {
		return errors.New("不正なコールです")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	if !(frame.FirstScore == 10 || frame.FirstScore+frame.SecondScore == 10) {
		return errors.New("不正なコールです")
	}
	frame.ThirdRestPins = currentPins
	if frame.SecondScore == 10 || frame.FirstScore+frame.SecondScore == 10 {
		frame.ThirdScore = 10 - len(frame.ThirdRestPins.GetRestNumbers())
	} else {
		frame.ThirdScore = len(frame.SecondRestPins.GetRestNumbers()) - len(frame.ThirdRestPins.GetRestNumbers())
	}
	frame.BowlCount++
	return nil
}
