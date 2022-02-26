package frame

import (
	"bowlingSystem/bowling/pins"
	"errors"
)

type Frame struct {
	FirstRestPins  *pins.Pins
	SecondRestPins *pins.Pins
	FirstScore     int
	SecondScore    int
	BowlCount      int
}

func New() *Frame {
	return &Frame{
		FirstRestPins:  nil,
		SecondRestPins: nil,
		FirstScore:     0,
		SecondScore:    0,
		BowlCount:      0,
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
