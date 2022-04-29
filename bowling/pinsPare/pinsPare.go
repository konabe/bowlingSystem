package pinsPare

import (
	"bowlingSystem/bowling/pins"
	"errors"
)

type PinsPair struct {
	// 何投目かのカウンター
	BowlCount int
	// １投目の残ピン
	FirstRestPins *pins.Pins
	// 2投目の残ピン
	SecondRestPins *pins.Pins
	FirstScore     int
	SecondScore    int
}

func New() *PinsPair {
	return &PinsPair{
		BowlCount:      0,
		FirstRestPins:  nil,
		SecondRestPins: nil,
		FirstScore:     0,
		SecondScore:    0,
	}
}

func (frame *PinsPair) BowlFirst(numbers []int) error {
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

func (frame *PinsPair) BowlSecond(numbers []int) error {
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