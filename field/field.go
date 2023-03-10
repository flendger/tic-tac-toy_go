package field

import "fmt"

type GameField struct {
	data [3][3]string
}

func (f *GameField) initField() {
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			f.data[i][y] = "_"
		}
	}
}

func (f *GameField) PrintField() {
	for _, line := range f.data {
		fmt.Println(line)
	}
	fmt.Println("**********")
}

func (f *GameField) SetItem(line uint8, col uint8, value string) error {
	currentValue := f.data[line-1][col-1]

	if currentValue != "_" {
		return fmt.Errorf("not empty")
	}

	f.data[line-1][col-1] = value
	return nil
}

func (f *GameField) CheckWin(value string) bool {
	//lines
	for line := 0; line < 3; line++ {
		var win = true
		for col := 0; col < 3; col++ {
			win = f.data[line][col] == value
			if !win {
				break
			}
		}

		if win {
			return true
		}
	}

	//cols
	for col := 0; col < 3; col++ {
		var win = true
		for line := 0; line < 3; line++ {
			win = f.data[line][col] == value
			if !win {
				break
			}
		}

		if win {
			return true
		}
	}

	//diagonal
	var win = true
	for x := 0; x < 3; x++ {
		win = f.data[x][x] == value
		if !win {
			break
		}
	}
	if win {
		return true
	}

	win = true
	for x := 0; x < 3; x++ {
		win = f.data[x][2-x] == value
		if !win {
			break
		}
	}

	return win
}

func (f *GameField) CheckDraw() bool {
	for line := 0; line < 3; line++ {
		for col := 0; col < 3; col++ {
			if f.data[line][col] == "_" {
				return false
			}
		}
	}

	return true
}

func CreateField() *GameField {
	newField := new(GameField)
	newField.initField()

	return newField
}
