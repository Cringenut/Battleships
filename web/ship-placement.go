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
	Size   int
	Coords []string
}

const size = 10

// Board represents the server board
var board [size][size]bool
var advancedBoard [size][size]bool

// ShipSizes defines the sizes of ships to be placed
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
		endCoords = []string{}
		firstCoord = PlacementCoordinate{}
		return
	}

	if currentPlacementPlacementType == data.Simple {
		placingShip = &ships[index]
	} else if currentPlacementPlacementType == data.Advanced {
		println("ADVANCED")
		placingShip = &advancedShips[index]
	}

	if currentPlacementPlacementType == data.Simple {
		// Clear previous ship's coordinates from the board if they exist
		if placingShip != nil && placingShip.Coords != nil {
			for _, coord := range placingShip.Coords {
				row, col, valid := getCoordPosition(strings.ToUpper(coord))
				if valid {
					board[row][col] = false
				}
			}
			placingShip.Coords = []string{}
		}
	} else if currentPlacementPlacementType == data.Advanced {
		// Clear previous ship's coordinates from the board if they exist
		if placingShip != nil && placingShip.Coords != nil {
			for _, coord := range placingShip.Coords {
				row, col, valid := getCoordPosition(strings.ToUpper(coord))
				if valid {
					advancedBoard[row][col] = false
				}
			}
			placingShip.Coords = []string{}
		}
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

	if placingShip.Size == 1 {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 1)
		SetLastCoord(coord)
	} else if placingShip != nil {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		if currentPlacementPlacementType == data.Simple || placingShip.Size < 3 {
			endCoords = possibleEndCoords(row, col, placingShip.Size)
		} else if currentPlacementPlacementType == data.Advanced {
			endCoords = possibleNextCoords(row, col)
			if len(endCoords) > 0 {
				placingShip.Coords = append(placingShip.Coords, coord)
			}
		}
	} else {
		firstCoord = PlacementCoordinate{}
	}
}

func possibleNextCoords(row, col int) []string {
	directions := []struct {
		x, y int
	}{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // Right, Down, Left, Up
	}
	var nextCoords []string

	for _, d := range directions {
		nextRow, nextCol := row+d.x, col+d.y
		if nextRow >= 0 && nextRow < size && nextCol >= 0 && nextCol < size && !advancedBoard[nextRow][nextCol] {
			nextCoords = append(nextCoords, fmt.Sprintf("%c%d", rune('A'+nextCol), data.InvertNumber(nextRow)))
		}
	}

	return nextCoords
}

func SetNextCoord(coord string) {
	_, _, valid := isValidPlacingCoordinate(strings.ToUpper(coord))

	if !valid {
		fmt.Println("Invalid coordinate. Please enter a valid coordinate.")
		firstCoord = PlacementCoordinate{}
		endCoords = []string{}
		return
	}

	placingShip.Coords = append(placingShip.Coords, coord)
	updateEndCoords()

	if len(endCoords) == 0 {
		fmt.Println("Choose another coordinate set")
		firstCoord = PlacementCoordinate{}
		endCoords = []string{}
		placingShip.Coords = placingShip.Coords[:len(placingShip.Coords)-1] // Remove the last invalid coord
		return
	}
}

func updateEndCoords() {
	var newEndCoords []string
	for _, shipCoord := range placingShip.Coords {
		row, col, valid := getCoordPosition(shipCoord)
		if valid {
			nextCoords := possibleNextCoords(row, col)
			for _, nextCoord := range nextCoords {
				if !data.StringSliceContains(placingShip.Coords, nextCoord) {
					newEndCoords = append(newEndCoords, nextCoord)
				}
			}
		}
	}
	endCoords = newEndCoords
}

func SetLastCoord(coord string) {
	if currentPlacementPlacementType == data.Simple {
		endRow, endCol, validEnd := isValidPlacingCoordinate(coord)

		if !validEnd || !data.StringSliceContains(endCoords, coord) {
			fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
			firstCoord = PlacementCoordinate{}
			endCoords = []string{}
			return
		}

		placeShip(firstCoord.Row, firstCoord.Col, endRow, endCol, placingShip.Size)
	} else if currentPlacementPlacementType == data.Advanced {
		if !data.StringSliceContains(endCoords, coord) {
			fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
			firstCoord = PlacementCoordinate{}
			endCoords = []string{}
			return
		}
		placingShip.Coords = append(placingShip.Coords, coord)
		placeAdvancedShip()
	}

	printBoard()
	firstCoord = PlacementCoordinate{}
	endCoords = []string{}
	placingShip = nil
}

// IsCoordinateInShips checks if the given coordinate string is inside any of the ships' coordinates
func IsCoordinateInShips(coord string) bool {
	if currentPlacementPlacementType == data.Simple {
		for _, ship := range ships {
			for _, shipCoord := range ship.Coords {
				if strings.EqualFold(shipCoord, coord) {
					return true
				}
			}
		}
	} else if currentPlacementPlacementType == data.Advanced {
		for _, advancedShip := range advancedShips {
			for _, shipCoord := range advancedShip.Coords {
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

	row, col, isValid := getCoordPosition(coord)

	if !isValid {
		return 0, 0, false
	}

	if currentPlacementPlacementType == data.Simple {
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
	} else {
		if advancedBoard[row][col] == true {
			return 0, 0, false
		}

		// Check adjacent cells, including diagonals
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				adjRow := row + dx
				adjCol := col + dy
				if adjRow >= 0 && adjRow < size && adjCol >= 0 && adjCol < size && advancedBoard[adjRow][adjCol] == true {
					return 0, 0, false
				}
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
			if currentPlacementPlacementType == data.Simple {
				if board[checkRow][checkCol] {
					valid = false
					break
				}
			} else {
				if advancedBoard[checkRow][checkCol] {
					valid = false
					break
				}
			}

			if currentPlacementPlacementType == data.Simple {
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
			} else {
				// Checking surrounding blocks to ensure no adjacent ships
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						adjRow, adjCol := checkRow+dx, checkCol+dy
						if adjRow >= 0 && adjRow < size && adjCol >= 0 && adjCol < size && advancedBoard[adjRow][adjCol] {
							valid = false
							break
						}
					}
					if !valid {
						break
					}
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

		// Add coordinate to the placingShip's Coords
		placingShip.Coords = append(placingShip.Coords, fmt.Sprintf("%c%d", rune('A'+curCol), 10-curRow))
	}
}

func placeAdvancedShip() {
	for _, coord := range placingShip.Coords {
		row, col, valid := getCoordPosition(coord)
		if valid {
			advancedBoard[row][col] = true
		}
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
	if currentPlacementPlacementType == data.Simple {
		if len(ships[index].Coords) == 0 {
			if &ships[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
		return strings.Join(ships[index].Coords, " ")
	} else if currentPlacementPlacementType == data.Advanced {
		println("ADVANCED")
		if len(advancedShips[index].Coords) == 0 {
			if &advancedShips[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
		return strings.Join(advancedShips[index].Coords, " ")
	}
	return ""
}

func ClearAllShipCoords() {
	if currentPlacementPlacementType == data.Simple {
		for i := range ships {
			ships[i].Coords = nil
		}
	} else if currentPlacementPlacementType == data.Advanced {
		println("ADVANCED")
		for i := range advancedShips {
			advancedShips[i].Coords = nil
		}
	}
}

func IsAnyShipMissingCoords() bool {
	if currentPlacementPlacementType == data.Simple {
		for i := range ships {
			if len(ships[i].Coords) != ships[i].Size {
				return false
			}
		}
	} else if currentPlacementPlacementType == data.Advanced {
		for i := range advancedShips {
			if len(advancedShips[i].Coords) != advancedShips[i].Size {
				return false
			}
		}
	}
	return true
}

func IsPlacingShipContains(coord string) bool {
	if placingShip == nil {
		return false
	}
	for _, shipCoord := range placingShip.Coords {
		if strings.EqualFold(shipCoord, coord) {
			return true
		}
	}
	return false
}
