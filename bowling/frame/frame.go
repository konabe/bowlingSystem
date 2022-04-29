package frame

import (
	"bowlingSystem/bowling/pins"
	"errors"
)

type Frame struct {
	// １投目の残ピン
	FirstRestPins *pins.Pins
	// 2投目の残ピン
	SecondRestPins *pins.Pins
	// 何投目かのカウンター
	BowlCount   int
	FirstScore  int
	SecondScore int
	// 10フレーム目のみの属性
	IsForTenFrame bool
	ThirdRestPins *pins.Pins
	ThirdScore    int
}

func New() *Frame {
	return &Frame{
		BowlCount:      0,
		FirstRestPins:  nil,
		SecondRestPins: nil,
		FirstScore:     0,
		SecondScore:    0,
		IsForTenFrame:  false,
		ThirdRestPins:  nil,
		ThirdScore:     0,
	}
}

func (frame *Frame) BowlFirst(numbers []int) error {
	newPins := pins.New()
	pins, err := newPins.KnockDown(numbers)
	if frame.BowlCount != 0 {
		return errors.New("不正なコールです")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	frame.FirstRestPins = pins
	frame.FirstScore = 10 - len(frame.FirstRestPins.GetRestNumbers())
	frame.BowlCount++
	return nil
}

func (frame *Frame) BowlSecond(numbers []int) error {
	if frame.IsForTenFrame {
		return frame.bowlSecondForTen(numbers)
	}
	return frame.bowlSecondForNormal(numbers)
}

func (frame *Frame) bowlSecondForNormal(numbers []int) error {
	pins, err := frame.FirstRestPins.KnockDown(numbers)
	if frame.BowlCount != 1 {
		return errors.New("不正なコールです")
	}
	if frame.FirstScore == 10 {
		return errors.New("すでにストライクをとっています")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	frame.SecondRestPins = pins
	frame.SecondScore = len(frame.FirstRestPins.GetRestNumbers()) - len(frame.SecondRestPins.GetRestNumbers())
	frame.BowlCount++
	return nil
}

func (frame *Frame) bowlSecondForTen(numbers []int) error {
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

func (frame *Frame) BowlThird(numbers []int) error {
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
