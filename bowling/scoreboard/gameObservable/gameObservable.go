package gameObservable

import "bowlingSystem/bowling/pinsPare"

type GameObservable interface {
	UpdateFrames([12]pinsPare.PinsPair)
}
