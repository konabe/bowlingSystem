package game

import (
	"bowlingSystem/bowling/pinsPare"
	"bowlingSystem/bowling/scoreboard/gameObservable"

	"errors"
)

type Game struct {
	Frames      [12]pinsPare.PinsPair
	FrameScores [12]int
	FrameIndex  int
	BowlCount   int
	Observable  gameObservable.GameObservable
}

func New() *Game {
	newGame := &Game{
		FrameIndex: 0,
		BowlCount:  0,
	}
	for i, _ := range newGame.Frames {
		newGame.Frames[i] = *pinsPare.New()
		newGame.FrameScores[i] = 0
	}
	return newGame
}

//Bowl ボウルを投げる
func (game *Game) Bowl(numbers []int) error {
	if !game.IsValidBowl() {
		return errors.New("Bowlメソッドが有効ではありません")
	}
	frame := game.Frames[game.FrameIndex]
	var err error
	if game.BowlCount == 0 {
		err = frame.BowlFirst(numbers)
		game.FrameScores[game.FrameIndex] = frame.FirstScore
	}
	if game.BowlCount == 1 {
		err = frame.BowlSecond(numbers)
		game.FrameScores[game.FrameIndex] += frame.SecondScore
	}
	game.Frames[game.FrameIndex] = frame
	game.Increment()
	if game.BowlCount == 1 && frame.FirstScore == 10 {
		game.Increment()
	}
	return err
}

//IsValidBowl Bowlメソッドのコールが有効かどうかを判定する
func (game *Game) IsValidBowl() bool {
	if game.FrameIndex < 0 || game.FrameIndex > len(game.Frames)-1 {
		return false
	}
	if game.BowlCount < 0 || game.BowlCount > 1 {
		return false
	}
	return true
}

//Increment 投球のインデックスを１つインクリメントする
func (game *Game) Increment() {
	if game.BowlCount == 0 {
		game.BowlCount++
		return
	}
	if game.FrameIndex < 12 && game.BowlCount == 1 {
		game.BowlCount = 0
		game.FrameIndex++
		return
	}
}

func (game Game) NotifyUpdate() {
	game.Observable.UpdateFrames(game.Frames)
}
