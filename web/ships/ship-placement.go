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
		tempShip := placingShip
		ClearData()
		placingShip = tempShip
		return
	}

	if placingShip.Size == 1 {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, 1)
		SetLastCoord(coord)
		return
	} else if placingShip != nil && data.GetCurrentPlacementPlacementType() == data.Simple {
		firstCoord = PlacementCoordinate{Row: row, Col: col, Coord: coord}
		endCoords = possibleEndCoords(row, col, placingShip.Size)
	} else if placingShip != nil && data.GetCurrentPlacementPlacementType() == data.Advanced {
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

	if !validEnd || !data.StringSliceContains(endCoords, coord) {
		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		ClearData()
		return
	}

	if data.GetCurrentPlacementPlacementType() == data.Simple || (data.GetCurrentPlacementPlacementType() == data.Advanced && placingShip.Size < 3) {
		placeShip(firstCoord.Row, firstCoord.Col, endRow, endCol, placingShip.Size)
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		placeShip(-1, -1, endRow, endCol, placingShip.Size)
	}

	printBoard()
	ClearData()
}

func SetNextCoord(coord string) {
	_, _, validNext := isValidPlacingCoordinate(coord)

	if !validNext || !data.StringSliceContains(endCoords, coord) {
		fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
		ClearData()
		return
	}

	endCoords = []string{}
	nextCoords = append(nextCoords, coord)

	for _, currentNextCoord := range nextCoords {
		row, col, _ := GetCoordPosition(currentNextCoord)
		for _, possibleEndCoord := range possibleEndCoords(row, col, 2) {
			if !data.StringSliceContains(nextCoords, possibleEndCoord) && possibleEndCoord != firstCoord.Coord {
				endCoords = append(endCoords, possibleEndCoord)
			}
		}
	}

	for _, possibleEndCoord := range possibleEndCoords(firstCoord.Row, firstCoord.Col, 2) {
		if !data.StringSliceContains(nextCoords, possibleEndCoord) {
			endCoords = append(endCoords, possibleEndCoord)
		}
	}

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
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		allCoords := nextCoords
		allCoords = append(allCoords, firstCoord.Coord)
		allCoords = append(allCoords, GetCoordString(endRow, endCol))

		fmt.Printf("All coords: %v", allCoords)
		for _, coord := range allCoords {
			placingShip.Coords = append(placingShip.Coords, coord)
			row, col, _ = GetCoordPosition(coord)
			board[row][col] = true
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
