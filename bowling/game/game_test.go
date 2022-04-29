package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 通常
func TestIsValidBowl(t *testing.T) {
	cases := []struct {
		frameIndex int
		bowlCount  int
		isValid    bool
	}{
		{frameIndex: -1, bowlCount: 0, isValid: false},
		{frameIndex: -1, bowlCount: 1, isValid: false},
		{frameIndex: -1, bowlCount: 2, isValid: false},
		{frameIndex: 0, bowlCount: -1, isValid: false},
		{frameIndex: 0, bowlCount: 0, isValid: true},
		{frameIndex: 0, bowlCount: 1, isValid: true},
		{frameIndex: 0, bowlCount: 2, isValid: false},
		{frameIndex: 0, bowlCount: 3, isValid: false},
		{frameIndex: 8, bowlCount: 0, isValid: true},
		{frameIndex: 8, bowlCount: 1, isValid: true},
		{frameIndex: 9, bowlCount: -1, isValid: false},
		{frameIndex: 9, bowlCount: 0, isValid: true},
		{frameIndex: 9, bowlCount: 1, isValid: true},
		{frameIndex: 9, bowlCount: 2, isValid: true},
		{frameIndex: 9, bowlCount: 3, isValid: false},
		{frameIndex: 10, bowlCount: 0, isValid: false},
		{frameIndex: 10, bowlCount: 1, isValid: false},
		{frameIndex: 10, bowlCount: 2, isValid: false},
	}
	for _, c := range cases {
		game := New()
		game.FrameIndex = c.frameIndex
		game.BowlCount = c.bowlCount
		msg := fmt.Sprintf("currentFrame: %d, bowlCount: %d", game.FrameIndex, game.BowlCount)
		assert.Equal(t, c.isValid, game.IsValidBowl(), msg)
	}
}

func TestIncrement(t *testing.T) {
	cases := []struct {
		frameIndex int
		bowlCount  int
	}{
		{frameIndex: 0, bowlCount: 0},
		{frameIndex: 0, bowlCount: 1},
		{frameIndex: 1, bowlCount: 0},
		{frameIndex: 1, bowlCount: 1},
		{frameIndex: 2, bowlCount: 0},
		{frameIndex: 2, bowlCount: 1},
		{frameIndex: 3, bowlCount: 0},
		{frameIndex: 3, bowlCount: 1},
		{frameIndex: 4, bowlCount: 0},
		{frameIndex: 4, bowlCount: 1},
		{frameIndex: 5, bowlCount: 0},
		{frameIndex: 5, bowlCount: 1},
		{frameIndex: 6, bowlCount: 0},
		{frameIndex: 6, bowlCount: 1},
		{frameIndex: 7, bowlCount: 0},
		{frameIndex: 7, bowlCount: 1},
		{frameIndex: 8, bowlCount: 0},
		{frameIndex: 8, bowlCount: 1},
		{frameIndex: 9, bowlCount: 0},
		{frameIndex: 9, bowlCount: 1},
		{frameIndex: 9, bowlCount: 2},
		{frameIndex: 9, bowlCount: 2},
	}
	game := New()
	for i, c := range cases {
		assert.Equal(t, c.frameIndex, game.FrameIndex, fmt.Sprintf("frameIndex: id=%d", i))
		assert.Equal(t, c.bowlCount, game.BowlCount, fmt.Sprintf("bowlCount: id=%d", i))
		game.Increment()
	}
}

func TestBowl(t *testing.T) {
	game := New()
	game.Bowl([]int{1, 2, 3})
	assert.Equal(t, [10]int{3, 0, 0, 0, 0, 0, 0, 0, 0, 0}, game.FrameScores)
	assert.Equal(t, true, game.IsValidBowl())
	game.Bowl([]int{4, 5, 6})
	assert.Equal(t, [10]int{6, 0, 0, 0, 0, 0, 0, 0, 0, 0}, game.FrameScores)
	game.Bowl([]int{4, 5, 6})
	game.Bowl([]int{})
	assert.Equal(t, [10]int{6, 3, 0, 0, 0, 0, 0, 0, 0, 0}, game.FrameScores)
	game.Bowl([]int{})
	game.Bowl([]int{})
	game.Bowl([]int{4, 5, 6})
	game.Bowl([]int{1, 2, 3, 7, 8, 9, 10})
	game.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game.Bowl([]int{})
	game.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, [10]int{6, 3, 0, 10, 10, 10, 0, 0, 0, 0}, game.FrameScores)
}
