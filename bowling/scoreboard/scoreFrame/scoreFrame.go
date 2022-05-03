package scoreFrame

import (
	"bowlingSystem/bowling/game/pins/pair"
	"bowlingSystem/bowling/scoreboard/score"
	"fmt"
)

type ScoreFrame struct {
	FirstScore      *score.Score
	SecondScore     *score.Score
	ThirdScore      *score.Score
	TotalScore      int
	CumulativeScore int
	IsDetermined    bool // ストライク等の計算が確定しているかどうか
	IsLastFrame     bool
}

func New() *ScoreFrame {
	return &ScoreFrame{
		FirstScore:   score.New(),
		SecondScore:  score.New(),
		ThirdScore:   score.New(),
		TotalScore:   0,
		IsDetermined: true,
		IsLastFrame:  false,
	}
}

func (scoreFrame ScoreFrame) IsStrike() bool {
	return scoreFrame.FirstScore.Symbol == score.Strike
}

func (scoreFrame ScoreFrame) IsSpare() bool {
	return scoreFrame.SecondScore.Symbol == score.Spare
}

func (scoreFrame ScoreFrame) Print(isCumulative bool) string {
	result := ""
	result += scoreFrame.FirstScore.Print()
	result += scoreFrame.SecondScore.Print()
	if scoreFrame.IsLastFrame {
		result += scoreFrame.ThirdScore.Print()
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

func (scoreFrame *ScoreFrame) UpdateCurrentFrame(currentFrame pair.PinsPair) {
	firstScore := currentFrame.FirstScore
	secondScore := currentFrame.SecondScore
	scoreFrame.FirstScore.Number = firstScore
	scoreFrame.FirstScore.Symbol = score.Unset
	scoreFrame.SecondScore.Number = secondScore
	scoreFrame.SecondScore.Symbol = score.Unset
	scoreFrame.TotalScore = firstScore + secondScore
	if firstScore == 10 {
		scoreFrame.FirstScore.Symbol = score.Strike
		scoreFrame.SecondScore.Symbol = score.Empty
		scoreFrame.IsDetermined = false
		return
	}
	if firstScore == 0 {
		scoreFrame.FirstScore.Symbol = score.Gutter
	}
	if firstScore+secondScore == 10 {
		scoreFrame.SecondScore.Symbol = score.Spare
		scoreFrame.IsDetermined = false
		return
	}
	if secondScore == 0 {
		scoreFrame.SecondScore.Symbol = score.Mistake
	}
}

func (scoreFrame *ScoreFrame) UpdateLastFrame(tensPinsPares []pair.PinsPair) {
	if len(tensPinsPares) != 3 {
		return
	}
	scoreFrame.IsLastFrame = true
	scoreFrame.FirstScore.Number = tensPinsPares[0].FirstScore
	scoreFrame.TotalScore = scoreFrame.FirstScore.Number
	scoreFrame.FirstScore.Symbol = score.Unset

	if tensPinsPares[0].FirstScore == 10 {
		scoreFrame.FirstScore.Symbol = score.Strike
		if tensPinsPares[1].FirstScore == 10 {
			scoreFrame.SecondScore.Number = 10
			scoreFrame.TotalScore += scoreFrame.SecondScore.Number
			scoreFrame.SecondScore.Symbol = score.Strike
			scoreFrame.ThirdScore.Number = tensPinsPares[2].FirstScore
			scoreFrame.TotalScore += scoreFrame.ThirdScore.Number
			scoreFrame.ThirdScore.Symbol = score.Unset
			if tensPinsPares[2].FirstScore == 10 {
				scoreFrame.ThirdScore.Symbol = score.Strike
			}
			if tensPinsPares[2].FirstScore == 0 {
				scoreFrame.ThirdScore.Symbol = score.Mistake
			}
			return
		}
		scoreFrame.SecondScore.Number = tensPinsPares[1].FirstScore
		scoreFrame.TotalScore += scoreFrame.SecondScore.Number
		scoreFrame.SecondScore.Symbol = score.Unset
		scoreFrame.ThirdScore.Number = tensPinsPares[1].SecondScore
		scoreFrame.TotalScore += scoreFrame.ThirdScore.Number
		scoreFrame.ThirdScore.Symbol = score.Unset
		if tensPinsPares[1].FirstScore == 0 {
			scoreFrame.SecondScore.Symbol = score.Mistake
		}
		if tensPinsPares[1].SecondScore == 0 {
			scoreFrame.ThirdScore.Symbol = score.Mistake
		}
		if tensPinsPares[1].FirstScore+tensPinsPares[1].SecondScore == 10 {
			scoreFrame.ThirdScore.Symbol = score.Spare
		}
		return
	}

	scoreFrame.SecondScore.Number = tensPinsPares[0].SecondScore
	scoreFrame.TotalScore += scoreFrame.SecondScore.Number
	scoreFrame.SecondScore.Symbol = score.Unset
	if tensPinsPares[0].FirstScore == 0 {
		scoreFrame.FirstScore.Symbol = score.Gutter
	}
	if tensPinsPares[0].FirstScore+tensPinsPares[0].SecondScore == 10 {
		scoreFrame.SecondScore.Symbol = score.Spare
		scoreFrame.ThirdScore.Number = tensPinsPares[1].FirstScore
		scoreFrame.TotalScore += scoreFrame.ThirdScore.Number
		scoreFrame.ThirdScore.Symbol = score.Unset
		if tensPinsPares[1].FirstScore == 10 {
			scoreFrame.ThirdScore.Symbol = score.Strike
		}
		if tensPinsPares[1].FirstScore == 0 {
			scoreFrame.ThirdScore.Symbol = score.Mistake
		}
		return
	}
	scoreFrame.ThirdScore.Symbol = score.Unset
	if tensPinsPares[0].SecondScore == 0 {
		scoreFrame.ThirdScore.Symbol = score.Mistake
	}
}
