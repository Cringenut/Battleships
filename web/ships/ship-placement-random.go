package ships

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	shipSizes = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
)

// Creating seed for random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Going slice with shipsSizes and try to place each one from slice
func GenerateRandomCoordinates() {
	// Reset the board and randomShips slice
	board = [10][10]bool{}
	randomShips = []string{}

	for _, size := range shipSizes {
		placeShipRandomly(size)
	}

	SetRandomShips(randomShips)
}

func placeShipRandomly(shipSize int) {
	// Try to place ship until successful
	for {
		// Creating random coordinate
		row, col, valid := isValidPlacingCoordinate(GetCoordString(rand.Intn(size), rand.Intn(size)))

		// Start placement if coordinate is valid and near space is empty
		if valid {
			// Using logic from placement
			// Setting first coordinate and find endCoords right after
			firstCoord = PlacementCoordinate{row, col, GetCoordString(row, col)}
			endCoords = possibleEndCoords(row, col, shipSize)

			// If any endCoord is available
			if len(endCoords) > 0 {
				// Choose random endCoord from the slice
				endCoord := endCoords[rand.Intn(len(endCoords))]
				endRow, endCol, _ := GetCoordPosition(endCoord)
				// Placing random ship
				placeRandomShip(firstCoord.Row, firstCoord.Col, endRow, endCol, shipSize)
				// Clearing data
				ClearData()
				// Stop trying to place the current ship
				break
			}
		}
	}
}

// Using the same logic to go from first coord to last and add coords between them as in Simple placement
func placeRandomShip(row, col, endRow, endCol int, length int) {
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

		randomShips = append(randomShips, fmt.Sprintf("%c%d", rune('A'+curCol), 10-curRow))
	}
}
