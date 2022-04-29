package gameObservable

import "bowlingSystem/bowling/frame"

type GameObservable interface {
	UpdateFrames([10]frame.Frame)
}
