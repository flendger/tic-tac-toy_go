package game

import (
	"fmt"
	"tic-tac-toy/field"
	"tic-tac-toy/human"
	"tic-tac-toy/robot"
)

type Game struct{}

func (receiver Game) Run(human *human.Human, robot *robot.Robot) {
	gameField := field.CreateField()

	for {
		for {
			humanLine, humanCol, err := human.Move()
			if err != nil {
				return
			}
			err = gameField.SetItem(humanLine, humanCol, "X")
			if err == nil {
				break
			}
		}

		if gameField.CheckWin("X") {
			gameField.PrintField()
			fmt.Println("Human win")
			break
		}

		if gameField.CheckDraw() {
			gameField.PrintField()
			fmt.Println("Draw!!!")
			break
		}

		for {
			robotLine, robotCol := robot.Move()
			err := gameField.SetItem(robotLine, robotCol, "O")
			if err == nil {
				break
			}
		}

		if gameField.CheckWin("O") {
			gameField.PrintField()
			fmt.Println("Bot win")
			break
		}

		if gameField.CheckDraw() {
			gameField.PrintField()
			fmt.Println("Draw!!!")
			break
		}

		gameField.PrintField()
	}
}
