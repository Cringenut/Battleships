package web

import (
	"Battleships/data"
	"fmt"
)

var placementTypes = []data.PlacementType{data.Simple, data.Advanced, data.Random}
var currentPlacementType data.PlacementType

func SetCurrentPlacementType(placementType data.PlacementType) {
	currentPlacementType = placementType
}

func GetCurrentPlacementType() data.PlacementType {
	return currentPlacementType
}

func SwitchCurrentPlacementType(isNext bool) {
	currentIndex := findPlacementIndex(currentPlacementType)
	if currentIndex == -1 {
		fmt.Println("Current placement type not found.")
		return
	}

	if isNext {
		if currentIndex+1 >= len(placementTypes) {
			currentPlacementType = placementTypes[0]
		} else {
			currentPlacementType = placementTypes[currentIndex+1]
		}
	} else {
		if currentIndex-1 < 0 {
			currentPlacementType = placementTypes[len(placementTypes)-1]
		} else {
			currentPlacementType = placementTypes[currentIndex-1]
		}
	}
}

func findPlacementIndex(value data.PlacementType) int {
	for i, v := range placementTypes {
		if v == value {
			return i
		}
	}
	return -1
}

func CanCurrentPlacementBeSaved() bool {
	switch currentPlacementType {
	case data.Simple:
		return IsAnyShipMissingCoords()
	default:
		return false
	}
}
