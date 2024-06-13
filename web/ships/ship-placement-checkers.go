package ships

import (
	"Battleships/data"
	"strings"
)

// isValidPlacingCoordinate checks if the coordinate is within board limits
func isValidPlacingCoordinate(coord string) (int, int, bool) {

	row, col, isValid := GetCoordPosition(coord)

	if !isValid {
		return -1, -1, false
	}

	// Ensure the spot is not already taken and check surrounding cells
	if board[row][col] == true {
		return -1, -1, false
	}

	// Check adjacent cells, including diagonals
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			adjRow := row + dx
			adjCol := col + dy
			if adjRow >= 0 && adjRow < size && adjCol >= 0 && adjCol < size && board[adjRow][adjCol] == true {
				return -1, -1, false
			}
		}
	}

	return row, col, true
}

// IsCoordinateInShips checks if the given coordinate string is inside any of the ships coordinates
func IsCoordinateInShips(coord string) bool {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for _, ship := range ships {
			for _, shipCoord := range ship.Coords {
				if strings.EqualFold(shipCoord, coord) {
					return true
				}
			}
		}
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
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

// Going through slices if ship length and amount of coordinates are different
// Return true meaning that ship missing coords
func IsAnyShipMissingCoords() bool {
	if data.GetCurrentPlacementPlacementType() == data.Simple {
		for i := range ships {
			if len(ships[i].Coords) != ships[i].Size {
				return true
			}
		}
		return false
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced {
		for i := range advancedShips {
			if len(advancedShips[i].Coords) != advancedShips[i].Size {
				return true
			}
		}
		return false
	}
	// For other cases check is not needed
	return false
}
