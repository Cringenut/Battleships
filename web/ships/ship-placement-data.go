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
var ships = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}

// Used in advanced placement
var advancedShips = []Ship{{4, nil}, {3, nil}, {3, nil}, {2, nil}, {2, nil},
	{2, nil}, {1, nil}, {1, nil}, {1, nil}, {1, nil}}

// Used in random placement
var randomShips []string

var placingShip *Ship
var firstCoord PlacementCoordinate
var endCoords []string
var nextCoords []string

func SetPlacingShip(index int) {
	// Fast way to clear placing ship and all according data
	if index < 0 {
		ClearData()
		return
	}

	// Choosing placingShip depending on current placement type
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

// Used in buttons to the right of placement board
func GetShipCoords(index int) string {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		// When placing ship it's coords are set to empty slice
		// Combined with comparing addresses can be found if this button is button of the ship we try to place
		if len(ships[index].Coords) == 0 {
			if &ships[index] == placingShip {
				return "Selected"
			}
			// If no coords are set return "+" showing that user has yet to place that ship
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

	// If coords are placed return them
	return strings.Join(nil, " ")
}

// Removing all ship coordinates from their slices
// Only used for Simple and Advanced placement types
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
	// Clearing board so new ship placement would be available
	ClearBoard()
}

// Getting ships coords depending on current placement placement type
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
		allCoords = GetRandomShips()
	} else {
		allCoords = []string{}
	}
	return allCoords
}

// Used when switching between placement types
func ClearBoard() {
	board = [size][size]bool{}
}

func ClearData() {
	placingShip = nil
	firstCoord = PlacementCoordinate{}
	endCoords = []string{}
	nextCoords = []string{}
}

// Filling the board depending on current placement
func RepopulateBoard() {
	for _, coord := range GetAllShipsCoords() {
		row, col, _ := GetCoordPosition(coord)
		board[row][col] = true
	}
}
