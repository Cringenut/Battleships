package web

import (
	"Battleships/data"
	"fmt"
)

// All available placement types in settings that can be switched
var placementTypes = []data.PlacementType{data.Simple, data.Advanced, data.Random}

// Initial placement style when settings is open or new placement type is chosen and saved
var currentSettingsPlacementType data.PlacementType

// Current placement when we choose the new ship formation
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

// Simple carosuel to switch between placement types if arrow button is clicked
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

// Default solution because golang doesn't provide standard library for that functionality
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
	// If any of the ships doesn't have all their coordinates
	case data.Simple:
		result = IsAnyShipMissingCoords()
	default:
		result = false
	}

	// Setting the currentSettingsPlacementType for the settings page
	if result == true {
		SetCurrentSettingsPlacementType(GetCurrentPlacementPlacementType())
	}
	return result
}
