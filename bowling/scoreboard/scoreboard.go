package scoreboard

import (
	"bowlingSystem/bowling/game/pins/pair"
	"bowlingSystem/bowling/scoreboard/score"
	"bowlingSystem/bowling/scoreboard/scoreFrame"
)

type Scoreboard struct {
	frames [10]scoreFrame.ScoreFrame
}

func New() *Scoreboard {
	scoreboard := &Scoreboard{}
	for i, _ := range scoreboard.frames {
		scoreboard.frames[i] = *scoreFrame.New()
	}
	return scoreboard
}

func (scoreboard Scoreboard) Print() string {
	result := ""
	for _, frame := range scoreboard.frames {
		result += frame.Print(true)
	}
	return result
}

func (scoreboard *Scoreboard) UpdateFrames(pinsPairs [12]pair.PinsPair) {
	for i, _ := range scoreboard.frames {
		if i < 9 {
			scoreboard.frames[i].UpdateCurrentFrame(pinsPairs[i])
		}
	}
	scoreboard.frames[9].UpdateLastFrame(pinsPairs[9:])
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
		if i >= 2 && scoreFrames[i-1].IsStrike() && scoreFrames[i-2].IsStrike() {
			scoreboard.frames[i-2].TotalScore += scoreFrame.FirstScore.Number
			scoreboard.frames[i-2].IsDetermined = true
		}
		if scoreFrames[i-1].IsSpare() {
			scoreboard.frames[i-1].TotalScore += scoreFrame.FirstScore.Number
			scoreboard.frames[i-1].IsDetermined = true
		}
		if scoreFrames[i-1].IsStrike() {
			scoreboard.frames[i-1].TotalScore += scoreFrame.FirstScore.Number + scoreFrame.SecondScore.Number
			if !scoreFrames[i].IsStrike() {
				scoreboard.frames[i-1].IsDetermined = true
			}
		}
		if i == 9 && scoreFrames[i].SecondScore.Symbol != score.Unset {
			scoreboard.frames[i-1].IsDetermined = true
		}
	}
}
