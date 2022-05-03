package gameObservable

import (
	"bowlingSystem/bowling/game/pins/pair"
)

type GameObservable interface {
	UpdateFrames([12]pair.PinsPair)
}
