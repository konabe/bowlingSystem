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
		Symbol: "",
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

func (scoreFrame ScoreFrame) print(isCumulative bool) string {
	result := ""
	result += scoreFrame.FirstScore.print()
	result += scoreFrame.SecondScore.print()
	if scoreFrame.IsLastFrame {
		result += scoreFrame.ThirdScore.print()
	}
	if scoreFrame.IsDetermined {
		if isCumulative {
			result += fmt.Sprintf("(%d)", scoreFrame.CumulativeScore)
		} else {
			result += fmt.Sprintf("(%d)", scoreFrame.TotalScore)
		}
	} else {
		result += "()"
	}
	return result
}

func (scoreFrame *ScoreFrame) updateCurrentFrame(currentFrame pinsPare.PinsPair) {
	firstScore := currentFrame.FirstScore
	secondScore := currentFrame.SecondScore
	scoreFrame.FirstScore.Number = firstScore
	scoreFrame.FirstScore.Symbol = Unset
	scoreFrame.SecondScore.Number = secondScore
	scoreFrame.SecondScore.Symbol = Unset
	scoreFrame.TotalScore = firstScore + secondScore
	if firstScore == 10 {
		scoreFrame.FirstScore.Symbol = Strike
		scoreFrame.SecondScore.Symbol = Empty
		scoreFrame.IsDetermined = false
		return
	}
	if firstScore == 0 {
		scoreFrame.FirstScore.Symbol = Gutter
	}
	if firstScore+secondScore == 10 {
		scoreFrame.SecondScore.Symbol = Spare
		scoreFrame.IsDetermined = false
		return
	}
	if secondScore == 0 {
		scoreFrame.SecondScore.Symbol = Mistake
	}
}

func (scoreFrame *ScoreFrame) updateLastFrame(tensPinsPares []pinsPare.PinsPair) {
	if len(tensPinsPares) != 3 {
		return
	}
	scoreFrame.IsLastFrame = true
	scoreFrame.FirstScore.Number = tensPinsPares[0].FirstScore
	scoreFrame.FirstScore.Symbol = Unset

	if tensPinsPares[0].FirstScore == 10 {
		scoreFrame.FirstScore.Symbol = Strike
		if tensPinsPares[1].FirstScore == 10 {
			scoreFrame.SecondScore.Number = 10
			scoreFrame.SecondScore.Symbol = Strike
			scoreFrame.ThirdScore.Number = tensPinsPares[2].FirstScore
			scoreFrame.ThirdScore.Symbol = Unset
			if tensPinsPares[2].FirstScore == 10 {
				scoreFrame.ThirdScore.Symbol = Strike
			}
			if tensPinsPares[2].FirstScore == 0 {
				scoreFrame.ThirdScore.Symbol = Mistake
			}
			return
		}
		scoreFrame.SecondScore.Number = tensPinsPares[1].FirstScore
		scoreFrame.SecondScore.Symbol = Unset
		scoreFrame.ThirdScore.Number = tensPinsPares[1].SecondScore
		scoreFrame.ThirdScore.Symbol = Unset
		if tensPinsPares[1].FirstScore == 0 {
			scoreFrame.SecondScore.Symbol = Mistake
		}
		if tensPinsPares[1].SecondScore == 0 {
			scoreFrame.ThirdScore.Symbol = Mistake
		}
		if tensPinsPares[1].FirstScore+tensPinsPares[1].SecondScore == 10 {
			scoreFrame.ThirdScore.Symbol = Spare
		}
		return
	}

	scoreFrame.SecondScore.Number = tensPinsPares[0].SecondScore
	scoreFrame.SecondScore.Symbol = Unset
	if tensPinsPares[0].FirstScore == 0 {
		scoreFrame.FirstScore.Symbol = Gutter
	}
	if tensPinsPares[0].FirstScore+tensPinsPares[0].SecondScore == 10 {
		scoreFrame.SecondScore.Symbol = Spare
		scoreFrame.ThirdScore.Number = tensPinsPares[1].FirstScore
		scoreFrame.ThirdScore.Symbol = Unset
		if tensPinsPares[1].FirstScore == 10 {
			scoreFrame.ThirdScore.Symbol = Strike
		}
		if tensPinsPares[1].FirstScore == 0 {
			scoreFrame.ThirdScore.Symbol = Mistake
		}
		return
	}
	scoreFrame.ThirdScore.Symbol = Unset
	if tensPinsPares[0].SecondScore == 0 {
		scoreFrame.ThirdScore.Symbol = Mistake
	}
}

type Scoreboard struct {
	frames [10]ScoreFrame
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
		result += frame.print(true)
	}
	return result
}

func (scoreboard *Scoreboard) UpdateFrames(pinsPairs [12]pinsPare.PinsPair) {
	for i, _ := range scoreboard.frames {
		if i < 9 {
			scoreboard.frames[i].updateCurrentFrame(pinsPairs[i])
		}
	}
	scoreboard.frames[9].updateLastFrame(pinsPairs[9:])
	scoreboard.frames[9].TotalScore =
		scoreboard.frames[9].FirstScore.Number +
			scoreboard.frames[9].SecondScore.Number +
			scoreboard.frames[9].ThirdScore.Number
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
		if i == 9 && scoreFrames[i].SecondScore.Symbol != Unset {
			scoreboard.frames[i-1].IsDetermined = true
		}
	}
}
