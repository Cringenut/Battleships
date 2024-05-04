package client

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
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func CalculateCellCoord(row int, col int) string {
	x := rune('A' + col)
	y := rune(InvertNumber(row))
	return string(x) + string(y)
}