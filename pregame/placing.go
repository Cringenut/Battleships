package pregame

import (
	"Battleships/data"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func PlaceShips() []string {
	scanner := bufio.NewScanner(os.Stdin)
	printBoard()

	for _, shipSize := range shipSizes {
		placed := false
		for !placed {
			fmt.Printf("Enter the start coordinate for a ship of length %d (e.g., A1): ", shipSize)
			scanner.Scan()
			input := strings.ToUpper(scanner.Text())
			row, col, valid := isValidCoordinate(input)

			if !valid {
				fmt.Println("Invalid coordinate. Please enter a valid coordinate.")
				continue
			}

			if shipSize > 1 {
				ends := possibleEndCoords(row, col, shipSize)
				if len(ends) == 0 {
					fmt.Println("No valid end positions from this start point. Please try a different start coordinate.")
					continue
				}

				fmt.Println("Possible end coordinates are:", ends)
				fmt.Print("Select an end coordinate from the list: ")
				scanner.Scan()
				endInput := strings.ToUpper(scanner.Text())
				endRow, endCol, validEnd := isValidCoordinate(endInput)

				if !validEnd || !data.StringSliceContains(ends, endInput) {
					fmt.Println("Invalid end coordinate. Please select a valid end coordinate from the list.")
					continue
				}

				placeShip(row, col, endRow, endCol, shipSize)

			} else {
				placeShip(row, col, row, col, shipSize)
			}

			printBoard()
			placed = true
		}
	}

	return getAllShipCoords()
}

const size = 10

// Board represents the server board
var board [size][size]bool

// ShipSizes defines the sizes of ships to be placed
var shipSizes = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

// isValidCoordinate checks if the coordinate is within board limits
func isValidCoordinate(coord string) (int, int, bool) {
	if len(coord) < 2 || !unicode.IsLetter(rune(coord[0])) || !unicode.IsDigit(rune(coord[1])) {
		return 0, 0, false
	}
	col := int(coord[0] - 'A')

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
		board[int(math.Min(float64(row), float64(endRow)))+i*dx][int(math.Min(float64(col), float64(endCol)))+i*dy] = true
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

func getAllShipCoords() []string {
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
