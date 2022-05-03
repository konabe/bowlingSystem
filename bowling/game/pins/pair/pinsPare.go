package pair

import (
	"bowlingSystem/bowling/game/pins"
	"errors"
)

type PinsPair struct {
	// 何投目かのカウンター
	bowlCount int
	// １投目の残ピン
	firstRestPins *pins.Pins
	// 2投目の残ピン
	secondRestPins *pins.Pins
	FirstScore     int
	SecondScore    int
}

func New() *PinsPair {
	return new(PinsPair)
}

func (frame *PinsPair) BowlFirst(numbers []int) error {
	newPins := pins.New()
	pins := newPins.KnockDown(numbers)
	if frame.bowlCount != 0 {
		return errors.New("不正なコールです")
	}
	frame.firstRestPins = pins
	frame.FirstScore = 10 - frame.firstRestPins.GetRestNumbers().Len()
	frame.bowlCount++
	return nil
}

func (frame *PinsPair) BowlSecond(numbers []int) error {
	pins := frame.firstRestPins.KnockDown(numbers)
	if frame.bowlCount != 1 {
		return errors.New("不正なコールです")
	}
	if frame.FirstScore == 10 {
		return errors.New("すでにストライクをとっています")
	}
	frame.secondRestPins = pins
	frame.SecondScore = frame.firstRestPins.GetRestNumbers().Len() - frame.secondRestPins.GetRestNumbers().Len()
	frame.bowlCount++
	return nil
}
