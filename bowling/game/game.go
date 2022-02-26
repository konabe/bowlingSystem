package game

import (
	"bowlingSystem/bowling/frame"
	"errors"
)

type Game struct {
	Frames        [9]frame.Frame
	FrameScores   [9]int
	TenFrame      frame.TenFrame
	TenFrameScore int
	FrameIndex    int
	BowlCount     int
}

func New() *Game {
	frames := [9]frame.Frame{}
	frameScores := [9]int{}
	for i := 0; i < len(frames); i++ {
		frames[i] = *frame.New()
		frameScores[i] = 0
	}

	return &Game{
		Frames:      frames,
		FrameScores: frameScores,
		FrameIndex:  0,
		BowlCount:   0,
	}
}

func (game *Game) Bowl(numbers []int) error {
	if !game.IsValidBowl() {
		return errors.New("Bowlメソッドが有効ではありません")
	}
	frame := game.Frames[game.FrameIndex]
	var err error
	if game.BowlCount == 0 {
		err = frame.BowlFirst(numbers)
	}
	if game.BowlCount == 1 {
		err = frame.BowlSecond(numbers)
	}
	game.Increment()
	return err
}

func (game *Game) IsValidBowl() bool {
	if game.FrameIndex < 0 || game.FrameIndex > 9 {
		return false
	}
	if game.BowlCount < 0 || game.BowlCount > 2 {
		return false
	}
	if game.FrameIndex != 9 && game.BowlCount > 1 {
		return false
	}
	return true
}

func (game *Game) Increment() {
	if game.BowlCount == 0 {
		game.BowlCount++
		return
	}
	if game.FrameIndex < 9 && game.BowlCount == 1 {
		game.BowlCount = 0
		game.FrameIndex++
		return
	}
	if game.FrameIndex == 9 && (game.BowlCount == 0 || game.BowlCount == 1) {
		game.BowlCount++
		return
	}
	if game.FrameIndex == 9 && game.BowlCount == 2 {
		return
	}
	if game.FrameIndex < 10 {
		return
	}
}
