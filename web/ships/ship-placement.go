package ships

import (
	"Battleships/data"
	"fmt"
	"math"
	"strings"
	"unicode"
)

type PlacementCoordinate struct {
	Row, Col int
	Coord    string
}

type Ship struct {
	size   int
	coords []string
}

const size = 10

// Board represents the server board and also to check for positions where ship can't be placed
var board [size][size]bool

// Used in default placement
// User
var ships = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}
var advancedShips = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}
var placingShip *Ship
var firstCoord PlacementCoordinate
var endCoords []string

func SetPlacingShip(index int) {
	if index < 0 {
		placingShip = nil
		firstCoord = PlacementCoordinate{}
		endCoords = []string{}
		return
	}

	if data.GetCurrentPlacementPlacementType() == data.Simple {
		placingShip = &ships[index]
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		println("ADVANCED")
		placingShip = &advancedShips[index]
	}

	// Clear previous ship's coordinates from the board if they exist
	if placingShip != nil && placingShip.coords != nil {
		for _, coord := range placingShip.coords {
			row, col, valid := GetCoordPosition(strings.ToUpper(coord))
			if valid {
				board[row][col] = false
			}
		}
		placingShip.coords = []string{}
	}
}

func GetPlacingShip() *Ship {
	return placingShip
}

func GetFirstCoord() PlacementCoordinate {
	return firstCoord
}

func GetEndCoords() []string {
	return endCoords
}

func SetFirstCoord(coord string) {
	row, col, valid := isValidPlacingCoordinate(strings.ToUpper(coord))

	if !valid {
		fmt.Println("Invalid coordinate. Please enter a valid coordinate.")
		firstCoord = PlacementCoordinate{}
		endCoords = []string{}
		return
	}

	if placingShip.size == 1 {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 1)
		SetLastCoord(coord)
	} else if placingShip != nil {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		if data.GetCurrentPlacementPlacementType() == data.Simple || placingShip.size < 3 {
			endCoords = possibleEndCoords(row, col, placingShip.size)
		} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
			endCoords = possibleEndCoords(row, col, 2)
		}
	} else {
		firstCoord = PlacementCoordinate{}
	}
}

func SetLastCoord(coord string) {
	endRow, endCol, validEnd := isValidPlacingCoordinate(coord)

	if !validEnd || !data.StringSliceContains(endCoords, coord) {
		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		firstCoord = PlacementCoordinate{}
		endCoords = []string{}
		return
	}

	placeShip(firstCoord.Row, firstCoord.Col, endRow, endCol, placingShip.size)
	printBoard()
	firstCoord = PlacementCoordinate{}
	endCoords = []string{}
	placingShip = nil
}

// IsCoordinateInShips checks if the given coordinate string is inside any of the ships' coordinates
func IsCoordinateInShips(coord string) bool {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for _, ship := range ships {
			for _, shipCoord := range ship.coords {
				if strings.EqualFold(shipCoord, coord) {
					return true
				}
			}
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for _, advancedShip := range advancedShips {
			for _, shipCoord := range advancedShip.coords {
				if strings.EqualFold(shipCoord, coord) {
					return true
				}
			}
		}
	}

	return false
}

// isValidPlacingCoordinate checks if the coordinate is within board limits
func isValidPlacingCoordinate(coord string) (int, int, bool) {

	row, col, isValid := GetCoordPosition(coord)

	if !isValid {
		return 0, 0, false
	}

	// Ensure the spot is not already taken and check surrounding cells
	if board[row][col] == true {
		return 0, 0, false
	}

	// Check adjacent cells, including diagonals
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			adjRow := row + dx
			adjCol := col + dy
			if adjRow >= 0 && adjRow < size && adjCol >= 0 && adjCol < size && board[adjRow][adjCol] == true {
				return 0, 0, false
			}
		}
	}

	return row, col, true
}

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
	return fmt.Sprintf("%c%d", 'A'+data.InvertNumber(row), col)
}

// possibleEndCoords finds possible placements for a ship of a given length from a start coordinate
func possibleEndCoords(row, col, length int) []string {
	directions := []struct {
		x, y int
	}{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // Right, Down, Left, Up
	}
	var ends []string

	for _, d := range directions {
		valid := true
		endRow, endCol := row+(length-1)*d.x, col+(length-1)*d.y
		if endRow < 0 || endRow >= size || endCol < 0 || endCol >= size {
			continue
		}
		for i := 0; i < length; i++ {
			checkRow, checkCol := row+i*d.x, col+i*d.y
			if board[checkRow][checkCol] {
				valid = false
				break
			}
			// Checking surrounding blocks to ensure no adjacent ships
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					adjRow, adjCol := checkRow+dx, checkCol+dy
					if adjRow >= 0 && adjRow < size && adjCol >= 0 && adjCol < size && board[adjRow][adjCol] {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
		}
		if valid {
			formattedEnd := fmt.Sprintf("%c%d", rune('A'+endCol), data.InvertNumber(endRow)) // Ensure correct formatting
			ends = append(ends, formattedEnd)
		}
	}
	return ends
}

func placeShip(row, col, endRow, endCol int, length int) {
	dx := 0
	dy := 0
	if row != endRow {
		dx = 1
	} else {
		dy = 1
	}
	for i := 0; i < length; i++ {
		curRow := int(math.Min(float64(row), float64(endRow))) + i*dx
		curCol := int(math.Min(float64(col), float64(endCol))) + i*dy
		board[curRow][curCol] = true

		// Add coordinate to the placingShip's coords
		placingShip.coords = append(placingShip.coords, fmt.Sprintf("%c%d", rune('A'+curCol), 10-curRow))
	}
}

func printBoard() {
	// Print each row
	for i := 0; i < len(board); i++ {
		// Print the row number right-aligned; 10-i to invert row order
		fmt.Printf("%2d ", 10-i)
		for _, cell := range board[i] {
			if cell {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

	// Print column labels
	fmt.Print("   ") // Add some initial spacing for row numbers
	for c := 'A'; c <= 'J'; c++ {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

}

func GetShipCoords(index int) string {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		if len(ships[index].coords) == 0 {
			if &ships[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		println("ADVANCED")
		if len(advancedShips[index].coords) == 0 {
			if &advancedShips[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
	}

	return strings.Join(ships[index].coords, " ")
}

func ClearAllShipCoords() {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for i := range ships {
			ships[i].coords = nil
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		println("ADVANCED")
		for i := range advancedShips {
			advancedShips[i].coords = nil
		}
	}
	ClearBoard()
}

func IsAnyShipMissingCoords() bool {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for i := range ships {
			if len(ships[i].coords) != ships[i].size {
				return true
			}
		}
		return false
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for i := range advancedShips {
			if len(advancedShips[i].coords) != advancedShips[i].size {
				return true
			}
		}
		return false
	}
	return true
}

func ClearBoard() {
	board = [size][size]bool{}
}

func GetShipsCoords() []string {
	var allCoords []string
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for _, ship := range ships {
			allCoords = append(allCoords, ship.coords...)
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for _, ship := range advancedShips {
			allCoords = append(allCoords, ship.coords...)
		}
	}
	return allCoords
}
