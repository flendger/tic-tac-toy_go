package robot

import "math/rand"

type Robot struct{}

func (r *Robot) Move() (line, col uint8) {

	line = uint8(rand.Intn(3)) + 1
	col = uint8(rand.Intn(3)) + 1

	return
}
