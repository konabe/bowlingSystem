package game

import (
	"bowlingSystem/bowling/scoreboard"
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
		{frameIndex: 9, bowlCount: 0, isValid: true},
		{frameIndex: 9, bowlCount: 1, isValid: true},
		{frameIndex: 10, bowlCount: 0, isValid: true},
		{frameIndex: 10, bowlCount: 1, isValid: true},
		{frameIndex: 11, bowlCount: 0, isValid: true},
		{frameIndex: 11, bowlCount: 1, isValid: true},
		{frameIndex: 12, bowlCount: 0, isValid: false},
		{frameIndex: 12, bowlCount: 1, isValid: false},
	}
	for _, c := range cases {
		game := New(nil)
		game.frameIndex = c.frameIndex
		game.bowlCount = c.bowlCount
		msg := fmt.Sprintf("currentFrame: %d, bowlCount: %d", game.frameIndex, game.bowlCount)
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
		{frameIndex: 10, bowlCount: 0},
		{frameIndex: 10, bowlCount: 1},
		{frameIndex: 11, bowlCount: 0},
		{frameIndex: 11, bowlCount: 1},
	}
	game := New(nil)
	for i, c := range cases {
		assert.Equal(t, c.frameIndex, game.frameIndex, fmt.Sprintf("frameIndex: id=%d", i))
		assert.Equal(t, c.bowlCount, game.bowlCount, fmt.Sprintf("bowlCount: id=%d", i))
		game.Increment()
	}
}

func TestBowl(t *testing.T) {
	game := New(nil)
	game.Bowl([]int{1, 2, 3})
	assert.Equal(t, 3, game.Pairs[0].FirstScore)
	assert.Equal(t, true, game.IsValidBowl())
	game.Bowl([]int{4, 5, 6})
	assert.Equal(t, 3, game.Pairs[0].FirstScore)
	assert.Equal(t, 3, game.Pairs[0].SecondScore)
	game.Bowl([]int{4, 5, 6})
	game.Bowl([]int{})
	assert.Equal(t, 3, game.Pairs[0].FirstScore)
	assert.Equal(t, 3, game.Pairs[0].SecondScore)
	assert.Equal(t, 3, game.Pairs[1].FirstScore)
	assert.Equal(t, 0, game.Pairs[1].SecondScore)
	game.Bowl([]int{})
	game.Bowl([]int{})
	assert.Equal(t, 0, game.Pairs[2].FirstScore)
	assert.Equal(t, 0, game.Pairs[2].SecondScore)
	game.Bowl([]int{4, 5, 6})
	game.Bowl([]int{1, 2, 3, 7, 8, 9, 10})
	assert.Equal(t, 3, game.Pairs[3].FirstScore)
	assert.Equal(t, 7, game.Pairs[3].SecondScore)
	game.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 10, game.Pairs[4].FirstScore)
	assert.Equal(t, 0, game.Pairs[4].SecondScore)
	game.Bowl([]int{})
	game.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 0, game.Pairs[5].FirstScore)
	assert.Equal(t, 10, game.Pairs[5].SecondScore)
}

func TestGame_NotifyUpdate(t *testing.T) {
	scoreboard1 := scoreboard.New()
	game := New(scoreboard1)
	game.Bowl([]int{1, 2, 3})
	game.Bowl([]int{4, 5})
	game.NotifyUpdate()
	assert.Equal(t, scoreboard1.Print(), "32(5)G-(5)G-(5)G-(5)G-(5)G-(5)G-(5)G-(5)G-(5)G0-(5)")
}
