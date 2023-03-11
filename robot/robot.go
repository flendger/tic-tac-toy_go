package robot

import (
	"math/rand"
	"tic-tac-toy/game/move"
)

type Robot struct{}

func (r *Robot) GetMarker() string {
	return "O"
}

func (r *Robot) GetName() string {
	return "Bot"
}

func (r *Robot) Move() (*move.Move, error) {
	line := uint8(rand.Intn(3)) + 1
	col := uint8(rand.Intn(3)) + 1

	return &move.Move{Line: line, Col: col}, nil
}
