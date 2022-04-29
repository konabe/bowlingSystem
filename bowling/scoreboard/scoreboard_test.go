package scoreboard

import (
	"bowlingSystem/bowling/frame"
	"bowlingSystem/bowling/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func bowlDummy() {

}

func TestCalculateScores(t *testing.T) {
	scoreboard := New()
	game1 := game.New()
	game1.Observable = scoreboard
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	game1.Bowl([]int{10})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Bowl([]int{})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7})
	game1.Bowl([]int{8, 9, 10})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	game1.Bowl([]int{10})
	game1.Bowl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	game1.Update()
	assert.Equal(t, "42(6)2-(8)", scoreboard.Print())
}

func TestUpdateCurrentFrame(t *testing.T) {
	cases := []struct {
		FirstNumbers  []int
		SecondNumbers []int
		ExpectedPrint string
	}{
		{[]int{1, 2, 3, 4}, []int{5, 6}, "42(6)"},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9, 10}, "4/()"},
		{[]int{}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "G/()"},
		{[]int{}, []int{}, "G-(0)"},
		{[]int{}, []int{1}, "G1(1)"},
		{[]int{1}, []int{}, "1-(1)"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{}, "X()"},
	}
	for _, c := range cases {
		scoreFrame := newScoreFrame()
		frame := frame.New()
		frame.BowlFirst(c.FirstNumbers)
		frame.BowlSecond(c.SecondNumbers)
		scoreFrame.updateCurrentFrame(*frame)
		assert.Equal(t, c.ExpectedPrint, scoreFrame.print())
	}
}
