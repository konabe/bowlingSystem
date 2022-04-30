package gameObservable

import "bowlingSystem/bowling/pinsPair"

type GameObservable interface {
	UpdateFrames([12]pinsPair.PinsPair)
}
