package player

import "tic-tac-toy/game/move"

type Player interface {
	GetMarker() string
	GetName() string
	Move() (*move.Move, error)
}
