package score

import "fmt"

type Symbol string

const (
	Strike  = Symbol("Strike")
	Spare   = Symbol("Spare")
	Mistake = Symbol("Mistake") // [ - ]
	Split   = Symbol("Split")
	Gutter  = Symbol("Gutter")
	Foul    = Symbol("Foul")
	Unset   = Symbol("Unset")
	Empty   = Symbol("Empty")
)

type Score struct {
	Number int
	Symbol Symbol
}

func New() *Score {
	return &Score{
		Number: 0,
		Symbol: "",
	}
}

func (score Score) Print() string {
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
