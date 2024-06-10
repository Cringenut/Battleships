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
	if col < 0 || col > 10 {
		return 0, 0, false
	}

	// Handle column parsing when it could be two digits (e.g., "10")
	row := 0
	if len(coord) == 3 && coord[1] == '1' && coord[2] == '0' {
		row = 0 // Zero-indexed, so '10' becomes 0
	} else if len(coord) == 2 && unicode.IsDigit(rune(coord[1])) {
		row = data.InvertNumber(int(coord[1] - '0'))
	} else {
		return 0, 0, false
	}

	if row < 0 || row >= size {
		return 0, 0, false
	}

	return row, col, true
}

// Convert row and column indices to a coordinate string
func GetCoordString(row, col int) string {
	return fmt.Sprintf("%c%d", 'A'+col, data.InvertNumber(row))
}
