package ships

import (
	"Battleships/data"
	"strings"
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

// Board represents the server board and also to check for positions where ship can't be placed
var board [size][size]bool

// Used in default placement
// User
var ships = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}
var advancedShips = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}
var randomShips []string
var placingShip *Ship
var firstCoord PlacementCoordinate
var endCoords []string
var nextCoords []string

func SetPlacingShip(index int) {
	if index < 0 {
		ClearData()
		return
	}

	if data.GetCurrentPlacementPlacementType() == data.Simple {
		placingShip = &ships[index]
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		placingShip = &advancedShips[index]
	}

	// Clear previous ship's coordinates from the board if they exist
	if placingShip != nil && placingShip.Coords != nil {
		for _, coord := range placingShip.Coords {
			row, col, valid := GetCoordPosition(strings.ToUpper(coord))
			if valid {
				board[row][col] = false
			}
		}
		placingShip.Coords = []string{}
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

func GetNextCoords() []string {
	return nextCoords
}

func GetRandomShips() []string {
	return randomShips
}

func SetRandomShips(ships []string) {
	randomShips = ships
}

func GetShipCoords(index int) string {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		if len(ships[index].Coords) == 0 {
			if &ships[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
		return strings.Join(ships[index].Coords, " ")
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		if len(advancedShips[index].Coords) == 0 {
			if &advancedShips[index] == placingShip {
				return "Selected"
			}
			return "+"
		}
		return strings.Join(advancedShips[index].Coords, " ")
	}

	return strings.Join(nil, " ")
}

func ClearAllShipsCoords() {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for i := range ships {
			ships[i].Coords = nil
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for i := range advancedShips {
			advancedShips[i].Coords = nil
		}
	}
	ClearBoard()
}

func GetAllShipsCoords() []string {
	var allCoords []string
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for _, ship := range ships {
			allCoords = append(allCoords, ship.Coords...)
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for _, ship := range advancedShips {
			allCoords = append(allCoords, ship.Coords...)
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Random {
		allCoords = randomShips
	} else {
		allCoords = []string{}
	}
	return allCoords
}

func ClearBoard() {
	board = [size][size]bool{}
}

func ClearData() {
	placingShip = nil
	firstCoord = PlacementCoordinate{}
	endCoords = []string{}
	nextCoords = []string{}
}

func RepopulateBoard() {
	for _, coord := range GetAllShipsCoords() {
		row, col, _ := GetCoordPosition(coord)
		board[row][col] = true
	}
}
