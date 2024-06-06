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

func CalculateEnemyAccuracy() float64 {
	if len(GetEnemyShots()) == len(GetGameStatus().OppShots) || len(GetGameStatus().OppShots) == 0 {
		return GetEnemyAccuracy()
	}

	hits := 0
	for _, shot := range GetGameStatus().OppShots {
		if StringSliceContains(GetPlayerShips(), shot) {
			hits++
		}
	}

	return (float64(hits) / float64(len(GetGameStatus().OppShots)) * 100)
}
