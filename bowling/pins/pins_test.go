package pins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRestNumbers(t *testing.T) {
	pins := New()
	pins, _ = pins.KnockDown([]int{1, 2, 3, 4})
	assert.Equal(t, []int{5, 6, 7, 8, 9, 10}, pins.GetRestNumbers())
}

func TestKnockDown(t *testing.T) {
	pins := New()
	newPins, err := pins.KnockDown([]int{1, 2, 3, 4, 5})
	assert.Equal(t, [10]bool{false, false, false, false, false, true, true, true, true, true}, newPins.Values)
	assert.Nil(t, err)
	newPinsNothing, errNothing := pins.KnockDown([]int{})
	assert.Equal(t, [10]bool{true, true, true, true, true, true, true, true, true, true}, newPinsNothing.Values)
	assert.Nil(t, errNothing)
	newPinsFull, errFull := pins.KnockDown([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, [10]bool{false, false, false, false, false, false, false, false, false, false}, newPinsFull.Values)
	assert.Nil(t, errFull)
	errorPins, errorErr := pins.KnockDown([]int{0})
	assert.Nil(t, errorPins)
	assert.NotNil(t, errorErr)
}

func TestKnockDownTwiceCalled(t *testing.T) {
	pins := New()
	firstPins, err := pins.KnockDown([]int{1, 2, 3})
	assert.Nil(t, err)
	SecondPins, err := firstPins.KnockDown([]int{4, 5, 6})
	assert.Nil(t, err)
	assert.Equal(t, []int{7, 8, 9, 10}, SecondPins.GetRestNumbers())
}

func TestIsValidNumbers(t *testing.T) {
	assert.True(t, IsValidNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	assert.True(t, IsValidNumbers([]int{1, 3, 5, 7, 9}))
	assert.True(t, IsValidNumbers([]int{1}))
	assert.True(t, IsValidNumbers([]int{})) // ピンが倒れないケースがあるので
	assert.False(t, IsValidNumbers([]int{0, 1, 3, 5, 7, 9}))
	assert.False(t, IsValidNumbers([]int{0}))
}
