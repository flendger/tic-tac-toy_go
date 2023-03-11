package human

import (
	"fmt"
	"tic-tac-toy/game/move"
	"tic-tac-toy/human/moveresolver"
)

type Human struct{}

func (r *Human) GetMarker() string {
	return "X"
}

func (r *Human) GetName() string {
	return "Human"
}

func (r *Human) Move() (*move.Move, error) {
	for {
		fmt.Print("Enter move [line col] (\\q - for exit): ")

		var (
			first  string
			second string
		)

		_, err := fmt.Scanln(&first, &second)
		if err != nil {
			if first == "\\q" {
				return nil, fmt.Errorf("exit")
			} else {
				continue
			}
		}

		line, col, err := moveresolver.Resolve(first, second)
		if err != nil {
			continue
		}

		return &move.Move{Line: line, Col: col}, nil
	}
}
