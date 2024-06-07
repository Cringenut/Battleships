package web

import (
	"Battleships/data"
	"fmt"
)

// Simple carosuel to switch between placement types if arrow button is clicked
func SwitchCurrentPlacementType(isNext bool) {
	currentIndex := findPlacementIndex(data.GetCurrentPlacementPlacementType())
	if currentIndex == -1 {
		fmt.Println("Current placement type not found.")
		return
	}

	if isNext {
		if currentIndex+1 >= len(data.GetPlacementTypes()) {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[0])
		} else {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[currentIndex+1])
		}
	} else {
		if currentIndex-1 < 0 {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[len(data.GetPlacementTypes())-1])
		} else {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[currentIndex-1])
		}
	}
}

// Default solution because golang doesn't provide standard library for that functionality
func findPlacementIndex(value data.PlacementType) int {
	for i, v := range data.GetPlacementTypes() {
		if v == value {
			return i
		}
	}
	return -1
}

func CanCurrentPlacementBeSaved() bool {
	result := false
	switch data.GetCurrentPlacementPlacementType() {
	// If any of the ships doesn't have all their coordinates
	case data.Simple:
		result = IsAnyShipNotMissingCoords()
	default:
		result = false
	}

	// Setting the currentSettingsPlacementType for the settings page
	if result == true {
		data.SetCurrentSettingsPlacementType(data.GetCurrentPlacementPlacementType())
	}
	return result
}
