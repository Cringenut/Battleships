package web

import (
	"Battleships/data"
	"fmt"
)

var placementTypes = []data.PlacementType{data.Simple, data.Advanced, data.Random}
var currentSettingsPlacementType data.PlacementType
var currentPlacementPlacementType data.PlacementType

func SetCurrentSettingsPlacementType(placementType data.PlacementType) {
	currentSettingsPlacementType = placementType
}

func GetCurrentSettingsPlacementType() data.PlacementType {
	return currentSettingsPlacementType
}

func SetCurrentPlacementPlacementType(placementType data.PlacementType) {
	currentPlacementPlacementType = placementType
}

func GetCurrentPlacementPlacementType() data.PlacementType {
	return currentPlacementPlacementType
}

func SwitchCurrentPlacementType(isNext bool) {
	currentIndex := findPlacementIndex(currentPlacementPlacementType)
	if currentIndex == -1 {
		fmt.Println("Current placement type not found.")
		return
	}

	if isNext {
		if currentIndex+1 >= len(placementTypes) {
			currentPlacementPlacementType = placementTypes[0]
		} else {
			currentPlacementPlacementType = placementTypes[currentIndex+1]
		}
	} else {
		if currentIndex-1 < 0 {
			currentPlacementPlacementType = placementTypes[len(placementTypes)-1]
		} else {
			currentPlacementPlacementType = placementTypes[currentIndex-1]
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
	result := false
	switch currentPlacementPlacementType {
	case data.Simple:
		result = IsAnyShipMissingCoords()
	default:
		result = false
	}

	if result == true {
		SetCurrentSettingsPlacementType(GetCurrentPlacementPlacementType())
	}
	return result
}
