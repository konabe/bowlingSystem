package pins

import (
	"bowlingSystem/bowling/game/pins/numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRestNumbers(t *testing.T) {
	pins := New()
	pins = pins.KnockDown([]int{1, 2, 3, 4})
	assert.Equal(t, pinNumbers.New(5, 6, 7, 8, 9, 10), pins.GetRestNumbers())
}

func TestKnockDown(t *testing.T) {
	pins := New()
	newPins := pins.KnockDown([]int{1, 2, 3, 4, 5})
	assert.Equal(t, [10]bool{false, false, false, false, false, true, true, true, true, true}, newPins.values)
	newPinsNothing := pins.KnockDown([]int{})
	assert.Equal(t, [10]bool{true, true, true, true, true, true, true, true, true, true}, newPinsNothing.values)
	newPinsFull := pins.KnockDown([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, [10]bool{false, false, false, false, false, false, false, false, false, false}, newPinsFull.values)
}

func TestKnockDownTwiceCalled(t *testing.T) {
	pins := New()
	firstPins := pins.KnockDown([]int{1, 2, 3})
	SecondPins := firstPins.KnockDown([]int{4, 5, 6})
	assert.Equal(t, pinNumbers.New(7, 8, 9, 10), SecondPins.GetRestNumbers())
}
