package human

import (
	"fmt"
	"tic-tac-toy/human/moveresolver"
)

type Human struct{}

func (r *Human) Move() (line, col uint8, err error) {

	for {
		fmt.Print("Enter move [line col] (\\q - for exit): ")

		var (
			first  string
			second string
		)

		_, err := fmt.Scanln(&first, &second)
		if err != nil {
			if first == "\\q" {
				return 0, 0, fmt.Errorf("exit")
			} else {
				continue
			}
		}

		line, col, err := moveresolver.Resolve(first, second)
		if err != nil {
			continue
		}

		return line, col, nil
	}
}
