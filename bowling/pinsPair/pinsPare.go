package pinsPair

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
	pins := newPins.KnockDown(numbers)
	if frame.BowlCount != 0 {
		return errors.New("不正なコールです")
	}
	frame.FirstRestPins = pins
	frame.FirstScore = 10 - frame.FirstRestPins.GetRestNumbers().Len()
	frame.BowlCount++
	return nil
}

func (frame *PinsPair) BowlSecond(numbers []int) error {
	pins := frame.FirstRestPins.KnockDown(numbers)
	if frame.BowlCount != 1 {
		return errors.New("不正なコールです")
	}
	if frame.FirstScore == 10 {
		return errors.New("すでにストライクをとっています")
	}
	frame.SecondRestPins = pins
	frame.SecondScore = frame.FirstRestPins.GetRestNumbers().Len() - frame.SecondRestPins.GetRestNumbers().Len()
	frame.BowlCount++
	return nil
}
