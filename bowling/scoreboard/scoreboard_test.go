package scoreboard

import (
	"bowlingSystem/bowling/game"
	"bowlingSystem/bowling/pinsPare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func bowlDummy(game *game.Game, first int, second int) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if first+second > 10 {
		return
	}
	game.Bowl(numbers[:first])
	if first != 10 {
		game.Bowl(numbers[first : first+second])
	}
}

func TestCalculateScores_min(t *testing.T) {
	scoreboard := New()
	game1 := game.New()
	game1.Observable = scoreboard
	for i := 0; i < 10; i++ {
		bowlDummy(game1, 0, 0)
	}
	game1.NotifyUpdate()
	assert.Equal(t, "G-(0)G-(0)G-(0)G-(0)G-(0)G-(0)G-(0)G-(0)G-(0)G0-(0)", scoreboard.Print())
}

func TestCalculateScores_max(t *testing.T) {
	scoreboard := New()
	game1 := game.New()
	game1.Observable = scoreboard
	for i := 0; i < 12; i++ {
		bowlDummy(game1, 10, 0)
	}
	game1.NotifyUpdate()
	assert.Equal(t, "X_(30)X_(60)X_(90)X_(120)X_(150)X_(180)X_(210)X_(240)X_(270)XXX(300)", scoreboard.Print())
}

func TestCalculateScores(t *testing.T) {
	scoreboard := New()
	game1 := game.New()
	game1.Observable = scoreboard
	bowlDummy(game1, 9, 1)
	bowlDummy(game1, 8, 0)
	bowlDummy(game1, 10, 0)
	bowlDummy(game1, 10, 0)
	bowlDummy(game1, 9, 0)
	bowlDummy(game1, 10, 0)
	bowlDummy(game1, 10, 0)
	bowlDummy(game1, 10, 0)
	bowlDummy(game1, 7, 3)
	bowlDummy(game1, 9, 1)
	bowlDummy(game1, 10, 0)
	game1.NotifyUpdate()
	assert.Equal(t, "9/(18)8-(26)X_(55)X_(74)9-(83)X_(113)X_(140)X_(160)7/(179)9/X(199)", scoreboard.Print())
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
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{}, "X_()"},
	}
	for _, c := range cases {
		scoreFrame := newScoreFrame()
		frame := pinsPare.New()
		frame.BowlFirst(c.FirstNumbers)
		frame.BowlSecond(c.SecondNumbers)
		scoreFrame.updateCurrentFrame(*frame)
		assert.Equal(t, c.ExpectedPrint, scoreFrame.print(false))
	}
}
