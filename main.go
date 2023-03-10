package main

import (
	"tic-tac-toy/game"
	"tic-tac-toy/human"
	"tic-tac-toy/robot"
)

func main() {
	humanPlayer := new(human.Human)
	robotPlayer := new(robot.Robot)

	gameInstance := new(game.Game)

	gameInstance.Run(humanPlayer, robotPlayer)
}
