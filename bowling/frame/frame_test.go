package frame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 通常
func TestBowlFirst(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, 3, frame.FirstScore)
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9, 10}, frame.FirstRestPins.GetRestNumbers())
	assert.Equal(t, 1, frame.bowlCount)
}

// 一投目がガーターのとき
func TestBowlFirst_Zero(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{})
	assert.Nil(t, err)
	assert.Equal(t, 0, frame.FirstScore)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, frame.FirstRestPins.GetRestNumbers())
	assert.Equal(t, 1, frame.bowlCount)
}

// 一投目でストライク
func TestBowlFirst_Full(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Nil(t, err)
	assert.Equal(t, 10, frame.FirstScore)
	assert.Equal(t, []int{}, frame.FirstRestPins.GetRestNumbers())
	assert.Equal(t, 1, frame.bowlCount)
}

// 異常系: 不正なコールのときにエラーになることをテストする
func TestBowl_Abuseful(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{})
	assert.Nil(t, err)
	err = frame.BowlFirst([]int{})
	assert.NotNil(t, err)
}
