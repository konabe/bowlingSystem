package scoreboard

import (
	"bowlingSystem/bowling/pinsPare"
	"fmt"
)

type ScoreSymbol string

const (
	Strike  = ScoreSymbol("Strike")
	Spare   = ScoreSymbol("Spare")
	Mistake = ScoreSymbol("Mistake") // [ - ]
	Split   = ScoreSymbol("Split")
	Gutter  = ScoreSymbol("Gutter")
	Foul    = ScoreSymbol("Foul")
	Unset   = ScoreSymbol("Unset")
	Empty   = ScoreSymbol("Empty")
)

type Score struct {
	Number int
	Symbol ScoreSymbol
}

func newScore() *Score {
	return &Score{
		Number: 0,
		Symbol: Unset,
	}
}

func (score Score) print() string {
	switch score.Symbol {
	case Strike:
		return "X"
	case Spare:
		return "/"
	case Mistake:
		return "-"
	case Split:
		return "S"
	case Gutter:
		return "G"
	case Foul:
		return "F"
	case Unset:
		return fmt.Sprint(score.Number)
	case Empty:
		return "_"
	default:
		return ""
	}
}

type ScoreFrame struct {
	FirstScore      *Score
	SecondScore     *Score
	ThirdScore      *Score
	TotalScore      int
	CumulativeScore int
	IsDetermined    bool // ストライク等の計算が確定しているかどうか
	IsLastFrame     bool
}

func newScoreFrame() *ScoreFrame {
	return &ScoreFrame{
		FirstScore:   newScore(),
		SecondScore:  newScore(),
		ThirdScore:   newScore(),
		TotalScore:   0,
		IsDetermined: true,
		IsLastFrame:  false,
	}
}

func (scoreFrame ScoreFrame) isStrike() bool {
	return scoreFrame.FirstScore.Symbol == Strike
}

func (scoreFrame ScoreFrame) isSpare() bool {
	return scoreFrame.SecondScore.Symbol == Spare
}

func (scoreFrame ScoreFrame) print() string {
	result := ""
	result += scoreFrame.FirstScore.print()
	result += scoreFrame.SecondScore.print()
	if scoreFrame.IsLastFrame {
		result += scoreFrame.ThirdScore.print()
	}
	if scoreFrame.IsDetermined {
		result += fmt.Sprintf("(%d)", scoreFrame.CumulativeScore)
	} else {
		result += "()"
	}
	return result
}

func (scoreFrame *ScoreFrame) updateCurrentFrame(currentFrame pinsPare.PinsPair) {
	firstScore := currentFrame.FirstScore
	secondScore := currentFrame.SecondScore
	if firstScore == 10 {
		scoreFrame.FirstScore.Number = firstScore
		scoreFrame.FirstScore.Symbol = Strike
		scoreFrame.SecondScore.Number = 0
		scoreFrame.SecondScore.Symbol = Empty
		scoreFrame.TotalScore = 10
		scoreFrame.IsDetermined = false
		return
	}
	if firstScore == 0 {
		scoreFrame.FirstScore.Symbol = Gutter
	}
	if firstScore+secondScore == 10 {
		scoreFrame.FirstScore.Number = firstScore
		scoreFrame.SecondScore.Number = secondScore
		scoreFrame.SecondScore.Symbol = Spare
		scoreFrame.TotalScore = 10
		scoreFrame.IsDetermined = false
		return
	}
	if secondScore == 0 {
		scoreFrame.SecondScore.Symbol = Mistake
	}
	scoreFrame.FirstScore.Number = firstScore
	scoreFrame.SecondScore.Number = secondScore
	scoreFrame.TotalScore = firstScore + secondScore
}

func (scoreFrame *ScoreFrame) updateLastFrame(lastFrame pinsPare.PinsPair) {
	scoreFrame.updateCurrentFrame(lastFrame)
	//thirdScore := scoreFrame.ThirdScore
}

type Scoreboard struct {
	frames [12]ScoreFrame
}

func New() *Scoreboard {
	scoreboard := &Scoreboard{}
	for i, _ := range scoreboard.frames {
		scoreboard.frames[i] = *newScoreFrame()
	}
	return scoreboard
}

func (scoreboard Scoreboard) Print() string {
	result := ""
	for _, frame := range scoreboard.frames {
		result += frame.print()
	}
	return result
}

func (scoreboard *Scoreboard) UpdateFrames(frames [12]pinsPare.PinsPair) {
	for i, _ := range scoreboard.frames {
		scoreboard.frames[i].updateCurrentFrame(frames[i])
	}
	scoreboard.recalculatePreviousFrames()
	for i, frame := range scoreboard.frames {
		if i == 0 {
			scoreboard.frames[i].CumulativeScore = frame.TotalScore
			continue
		}
		scoreboard.frames[i].CumulativeScore = frame.TotalScore + scoreboard.frames[i-1].CumulativeScore
	}
}

func (scoreboard *Scoreboard) recalculatePreviousFrames() {
	scoreFrames := scoreboard.frames
	for i, scoreFrame := range scoreFrames {
		if i == 0 {
			continue
		}
		if i >= 2 && scoreFrames[i-1].isStrike() && scoreFrames[i-2].isStrike() {
			scoreboard.frames[i-2].TotalScore += scoreFrame.FirstScore.Number
			scoreboard.frames[i-2].IsDetermined = true
		}
		if scoreFrames[i-1].isSpare() {
			scoreboard.frames[i-1].TotalScore += scoreFrame.FirstScore.Number
			scoreboard.frames[i-1].IsDetermined = true
		}
		if scoreFrames[i-1].isStrike() {
			scoreboard.frames[i-1].TotalScore += scoreFrame.FirstScore.Number + scoreFrame.SecondScore.Number
			if !scoreFrames[i].isStrike() {
				scoreboard.frames[i-1].IsDetermined = true
			}
		}
	}
}
