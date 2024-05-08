package data

import (
	"strconv"
)

func InvertNumber(digit int) int {
	switch digit {
	case 1:
		return 9
	case 2:
		return 8
	case 3:
		return 7
	case 4:
		return 6
	case 5:
		return 5
	case 6:
		return 4
	case 7:
		return 3
	case 8:
		return 2
	case 9:
		return 1
	default:
		return 10
	}
}

// Helper function to check if a string is in a slice of strings
func StringSliceContains(s []string, str string) bool {
	if str[0] == 'p' {
		str = str[1:]
	}
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func CalculateCellCoord(row int, col int) string {
	x := InvertNumber(row)
	y := rune('A' + col)
	return string(y) + strconv.Itoa(x)
}

func IsCellHit(coord string) (bool, bool) {
	hit, ok := GetPlayerShots()[coord]
	if ok {
		return true, hit
	} else {
		return false, false
	}
}
