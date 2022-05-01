package pins

import (
	pinNumbers "bowlingSystem/bowling/pins/numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRestNumbers(t *testing.T) {
	pins := New()
	pins, _ = pins.KnockDown([]int{1, 2, 3, 4})
	assert.Equal(t, pinNumbers.New(5, 6, 7, 8, 9, 10), pins.GetRestNumbers())
}

func TestKnockDown(t *testing.T) {
	pins := New()
	newPins, err := pins.KnockDown([]int{1, 2, 3, 4, 5})
	assert.Equal(t, [10]bool{false, false, false, false, false, true, true, true, true, true}, newPins.values)
	assert.Nil(t, err)
	newPinsNothing, errNothing := pins.KnockDown([]int{})
	assert.Equal(t, [10]bool{true, true, true, true, true, true, true, true, true, true}, newPinsNothing.values)
	assert.Nil(t, errNothing)
	newPinsFull, errFull := pins.KnockDown([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, [10]bool{false, false, false, false, false, false, false, false, false, false}, newPinsFull.values)
	assert.Nil(t, errFull)
}

func TestKnockDownTwiceCalled(t *testing.T) {
	pins := New()
	firstPins, err := pins.KnockDown([]int{1, 2, 3})
	assert.Nil(t, err)
	SecondPins, err := firstPins.KnockDown([]int{4, 5, 6})
	assert.Nil(t, err)
	assert.Equal(t, pinNumbers.New(7, 8, 9, 10), SecondPins.GetRestNumbers())
}
