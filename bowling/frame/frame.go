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
	bowlCount      int
}

func New() *Frame {
	return &Frame{
		FirstRestPins:  pins.New(),
		SecondRestPins: pins.New(),
		FirstScore:     0,
		SecondScore:    0,
		bowlCount:      0,
	}
}

func (frame *Frame) BowlFirst(numbers []int) error {
	pins, err := frame.FirstRestPins.KnockDown(numbers)
	if frame.bowlCount != 0 {
		return errors.New("不正なコールです")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	frame.FirstRestPins = pins
	frame.FirstScore = 10 - len(frame.FirstRestPins.GetRestNumbers())
	frame.bowlCount++
	return nil
}

func (frame *Frame) ThrowSecond(numbers []int) error {
	pins, err := frame.FirstRestPins.KnockDown(numbers)
	if frame.bowlCount != 1 {
		return errors.New("不正なコールです")
	}
	if err != nil {
		return errors.New("numbersが有効ではありません")
	}
	frame.SecondRestPins = pins
	frame.SecondScore = frame.FirstScore - len(frame.SecondRestPins.GetRestNumbers())
	frame.bowlCount++
	return nil
}
