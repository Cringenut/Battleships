package ships

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	boardSize = 10
)

var (
	shipSizes = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomCoordinates() {
	for _, size := range shipSizes {
		placeShipRandomly(size)
	}

	SetRandomShips(randomShips)
}

func placeShipRandomly(shipSize int) {
	for {
		row, col, valid := isValidPlacingCoordinate(GetCoordString(rand.Intn(boardSize), rand.Intn(boardSize)))

		if valid {
			firstCoord = PlacementCoordinate{row, col, GetCoordString(row, col)}
			endCoords = possibleEndCoords(row, col, shipSize)

			if len(endCoords) > 0 {
				endCoord := endCoords[rand.Intn(len(endCoords))]
				endRow, endCol, _ := GetCoordPosition(endCoord)
				placeRandomShip(firstCoord.Row, firstCoord.Col, endRow, endCol, shipSize)
				ClearData()
				break
			}
		}
	}
}

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
