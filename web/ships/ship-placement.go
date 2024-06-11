package ships

import (
	"Battleships/data"
	"fmt"
	"math"
	"strings"
)

func SetFirstCoord(coord string) {
	row, col, valid := isValidPlacingCoordinate(strings.ToUpper(coord))

	if !valid {
		fmt.Println("Invalid coordinate. Please enter a valid coordinate.")
		// Used to keep trying to place the same ship
		// Because clear data clears placingShip too
		tempShip := placingShip
		ClearData()
		placingShip = tempShip
		return
	}

	// If ship is of size 1 try place it on the board
	if placingShip.Size == 1 {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 1)
		SetLastCoord(coord)
		return
	} else if placingShip != nil && data.GetCurrentPlacementPlacementType() == data.Simple {
		// Simple ship placement
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, placingShip.Size)
	} else if placingShip != nil && data.GetCurrentPlacementPlacementType() == data.Advanced {
		// Advanced ship placement
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 2)
	} else {
		firstCoord = PlacementCoordinate{}
	}

	if len(endCoords) == 0 {
		fmt.Println("Invalid ship placement. Please choose different coordinate.")
		ClearData()
	}
}

func SetLastCoord(coord string) {
	endRow, endCol, validEnd := isValidPlacingCoordinate(coord)

	// Checking if coordinate is within board and among endCoords
	if !validEnd || !data.StringSliceContains(endCoords, coord) {

		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		// Used to keep trying to place the same ship
		// Because clear data clears placingShip too
		tempShip := placingShip
		ClearData()
		placingShip = tempShip
		return
	}

	// The second condition is used for ships of length 2 because they can use the same logic as in Simple placement
	if data.GetCurrentPlacementPlacementType() == data.Simple || (data.GetCurrentPlacementPlacementType() == data.Advanced && placingShip.Size < 3) {
		placeShip(firstCoord.Row, firstCoord.Col, endRow, endCol, placingShip.Size)
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		// row and col are not required for advanced ship placement
		placeShip(-1, -1, endRow, endCol, placingShip.Size)
	}

	// Debug
	printBoard()
	// Stop placing the ship
	ClearData()
}

// Advanced placement if not placing last coord
func SetNextCoord(coord string) {
	// Checking if given coord is within board
	_, _, validNext := isValidPlacingCoordinate(coord)

	// Checking if coordinate is within board and among endCoords
	if !validNext || !data.StringSliceContains(endCoords, coord) {
		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		ClearData()
		return
	}

	// Clearing coords to go through first and found next coords later
	endCoords = []string{}
	nextCoords = append(nextCoords, coord)

	// Finding available next or last coordinates
	for _, currentNextCoord := range nextCoords {
		row, col, _ := GetCoordPosition(currentNextCoord)
		// Checking for the ship of size 2
		// It checks 1 cell next to out currentNextCoord
		for _, possibleEndCoord := range possibleEndCoords(row, col, 2) {
			// Append only if coordinate wasn't found yet and not first coord
			if !data.StringSliceContains(nextCoords, possibleEndCoord) && possibleEndCoord != firstCoord.Coord {
				endCoords = append(endCoords, possibleEndCoord)
			}
		}
	}

	// Appending the rest of available end coordinates including next to the first one
	for _, possibleEndCoord := range possibleEndCoords(firstCoord.Row, firstCoord.Col, 2) {
		if !data.StringSliceContains(nextCoords, possibleEndCoord) {
			endCoords = append(endCoords, possibleEndCoord)
		}
	}

	// If no space left to place clear current ship placement
	if len(endCoords) == 0 {
		fmt.Println("Invalid ship placement. Please choose different coordinate.")
		ClearData()
	}
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
	if data.GetCurrentPlacementPlacementType() == data.Simple || placingShip.Size == 1 {
		// Direction depending on ship orientation
		dx := 0
		dy := 0
		if row != endRow {
			dx = 1
		} else {
			dy = 1
		}
		for i := 0; i < length; i++ {
			// Going from first coord to end coord direction
			curRow := int(math.Min(float64(row), float64(endRow))) + i*dx
			curCol := int(math.Min(float64(col), float64(endCol))) + i*dy
			// Populating the board
			board[curRow][curCol] = true

			// Add coordinate to the placingShip's Coords
			placingShip.Coords = append(placingShip.Coords, fmt.Sprintf("%c%d", rune('A'+curCol), 10-curRow))
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		// Combining first, nextCoords, and endCoord
		allCoords := nextCoords
		allCoords = append(allCoords, firstCoord.Coord)
		allCoords = append(allCoords, GetCoordString(endRow, endCol))

		fmt.Printf("All coords: %v", allCoords)
		// Simply append all coord to the ship
		for _, coord := range allCoords {
			placingShip.Coords = append(placingShip.Coords, coord)
			row, col, _ = GetCoordPosition(coord)
			// Populating the board
			board[row][col] = true
		}
	}

}

// Print board for debug
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
