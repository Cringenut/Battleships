package data

// All available placement types in settings that can be switched
var placementTypes = []PlacementType{Simple, Advanced, Random}

// Initial placement style when settings is open or new placement type is chosen and saved
var currentSettingsPlacementType PlacementType

// Current placement when we choose the new ship formation
var currentPlacementPlacementType PlacementType

func SetCurrentSettingsPlacementType(placementType PlacementType) {
	currentSettingsPlacementType = placementType
}

func GetCurrentSettingsPlacementType() PlacementType {
	return currentSettingsPlacementType
}

func SetCurrentPlacementPlacementType(placementType PlacementType) {
	currentPlacementPlacementType = placementType
}

func GetCurrentPlacementPlacementType() PlacementType {
	return currentPlacementPlacementType
}

func GetPlacementTypes() []PlacementType {
	return placementTypes
}
