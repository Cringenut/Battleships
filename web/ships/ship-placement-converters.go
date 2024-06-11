package ships

import (
	"Battleships/data"
	"fmt"
	"unicode"
)

func GetCoordPosition(coord string) (int, int, bool) {
	if len(coord) < 2 || !unicode.IsLetter(rune(coord[0])) || !unicode.IsDigit(rune(coord[1])) {
		return 0, 0, false
	}
	col := int(coord[0] - 'A')
	// If letter goes beyond "J" return not valid coord
	if col < 0 || col > 10 {
		return 0, 0, false
	}

	// Handle column parsing when it could be two digits, in our case 10
	row := 0
	if len(coord) == 3 && coord[1] == '1' && coord[2] == '0' {
		// Row is zero because our board is reversed and starts with 10
		row = 0
		// If coord is valid and has letter with digit after it
	} else if len(coord) == 2 && unicode.IsDigit(rune(coord[1])) {
		row = data.InvertNumber(int(coord[1] - '0'))
	} else {
		return 0, 0, false
	}

	// If not in range [1:10]
	if row < 0 || row >= size {
		return 0, 0, false
	}

	return row, col, true
}

// Convert row and column indices to a coordinate string
func GetCoordString(row, col int) string {
	return fmt.Sprintf("%c%d", 'A'+col, data.InvertNumber(row))
}
