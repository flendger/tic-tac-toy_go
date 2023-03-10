package moveresolver

import (
	"fmt"
	"strconv"
)

func Resolve(first string, second string) (uint8, uint8, error) {
	var line uint64
	var col uint64

	line, err := strconv.ParseUint(first, 10, 8)
	if err != nil {
		return 0, 0, err
	}

	col, err = strconv.ParseUint(second, 10, 8)
	if err != nil {
		return 0, 0, err
	}

	if outOfRange(line) || outOfRange(col) {
		return 0, 0, fmt.Errorf("incorrect coords")
	}

	return uint8(line), uint8(col), nil
}

func outOfRange(position uint64) bool {
	return position < 1 || position > 3
}
