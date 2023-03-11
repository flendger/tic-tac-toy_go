package game

import (
	"fmt"
	"tic-tac-toy/field"
	"tic-tac-toy/game/player"
)

type Game struct{}

func (g *Game) Run(players ...player.Player) {
	gameField := field.New()

	for {
		for _, currentPlayer := range players {
			for {
				move, err := currentPlayer.Move()
				if err != nil {
					return
				}

				moveErr := gameField.SetItem(move.Line, move.Col, currentPlayer.GetMarker())
				if moveErr == nil {
					break
				}
			}

			if gameField.CheckWin(currentPlayer.GetMarker()) {
				gameField.PrintField()
				fmt.Println(currentPlayer.GetName(), "win")
				return
			}

			if gameField.CheckDraw() {
				gameField.PrintField()
				fmt.Println("Draw!!!")
				return
			}

			gameField.PrintField()
		}
	}
}
