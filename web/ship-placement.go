package web

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

// Board represents the server board
var board [size][size]bool

// ShipSizes defines the sizes of ships to be placed
var ships = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}
var placingShip *Ship
var firstCoord PlacementCoordinate
var endCoords []string

func SetPlacingShip(index int) {
	placingShip = &ships[index]

	// Clear previous ship's coordinates from the board if they exist
	if placingShip != nil && placingShip.coords != nil {
		for _, coord := range placingShip.coords {
			row, col, valid := getCoordPosition(strings.ToUpper(coord))
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
		return
	}

	if placingShip.size == 1 {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 1)
		SetLastCoord(coord)
	} else if placingShip != nil {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, placingShip.size)
	} else {
		firstCoord = PlacementCoordinate{}
	}
}

func SetLastCoord(coord string) {
	endRow, endCol, validEnd := isValidPlacingCoordinate(coord)

	if !validEnd || !data.StringSliceContains(endCoords, coord) {
		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		firstCoord = PlacementCoordinate{}
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
	for _, ship := range ships {
		for _, shipCoord := range ship.coords {
			if strings.EqualFold(shipCoord, coord) {
				return true
			}
		}
	}
	return false
}

// isValidPlacingCoordinate checks if the coordinate is within board limits
func isValidPlacingCoordinate(coord string) (int, int, bool) {

	row, col, isValid := getCoordPosition(coord)

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

func getCoordPosition(coord string) (int, int, bool) {
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

func GetAllShipCoords() []string {
	var coords []string
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] {
				// Convert board indices back to user-friendly coordinates
				coord := fmt.Sprintf("%c%d", rune('A'+col), 10-row) // Adjust for zero-index and reverse row order
				coords = append(coords, coord)
			}
		}
	}
	return coords
}

func GetShipCoords(index int) string {
	if len(ships[index].coords) == 0 {
		return "+"
	}
	return strings.Join(ships[index].coords, " ")
}
