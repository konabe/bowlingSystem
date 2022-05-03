package game

import (
	"bowlingSystem/bowling/game/pins/pair"
	"bowlingSystem/bowling/scoreboard/gameObservable"

	"errors"
)

type Game struct {
	observable gameObservable.GameObservable

	Pairs      [12]pair.PinsPair
	frameIndex int
	bowlCount  int
}

func New(observable gameObservable.GameObservable) *Game {
	newGame := new(Game)
	newGame.observable = observable
	for i, _ := range newGame.Pairs {
		newGame.Pairs[i] = *pair.New()
	}
	return newGame
}

//Bowl ボウルを投げる
func (game *Game) Bowl(numbers []int) error {
	if !game.IsValidBowl() {
		return errors.New("bowlメソッドが有効ではありません")
	}
	frame := game.Pairs[game.frameIndex]
	var err error
	if game.bowlCount == 0 {
		err = frame.BowlFirst(numbers)
	}
	if game.bowlCount == 1 {
		err = frame.BowlSecond(numbers)
	}
	game.Pairs[game.frameIndex] = frame
	game.Increment()
	if game.bowlCount == 1 && frame.FirstScore == 10 {
		game.Increment()
	}
	return err
}

//IsValidBowl Bowlメソッドのコールが有効かどうかを判定する
func (game *Game) IsValidBowl() bool {
	if game.frameIndex < 0 || game.frameIndex > len(game.Pairs)-1 {
		return false
	}
	if game.bowlCount < 0 || game.bowlCount > 1 {
		return false
	}
	return true
}

//Increment 投球のインデックスを１つインクリメントする
func (game *Game) Increment() {
	if game.bowlCount == 0 {
		game.bowlCount = 1
		return
	}
	if game.frameIndex < 12 && game.bowlCount == 1 {
		game.bowlCount = 0
		game.frameIndex++
		return
	}
}

func (game Game) NotifyUpdate() {
	game.observable.UpdateFrames(game.Pairs)
}
