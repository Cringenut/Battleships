package web

import (
	"Battleships/data"
	"Battleships/web/ships"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func SaveSettings(c *gin.Context) string {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to read request body")
		return ""
	}

	// Parse the form data
	formData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to parse form data")
		return ""
	}

	saveNickname := formData.Get("nickname")
	saveDescription := formData.Get("description")

	fmt.Println("NICKNAME: " + saveNickname)
	fmt.Println("DESCRIPTION: " + saveDescription)

	if saveNickname == "" {
		return saveDescription
	}

	data.SetPlayerData(saveNickname, saveDescription)
	if ships.IsAnyShipMissingCoords() {
		data.SetPlayerShipPlacementType(data.GetCurrentSettingsPlacementType())
		data.SetPlayerShips(ships.GetShipsCoords())
	}

	Redirect(c, "/")
	return saveDescription
}

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
		result = !ships.IsAnyShipMissingCoords()
	default:
		result = false
	}

	// Setting the currentSettingsPlacementType for the settings page
	if result == true {
		data.SetCurrentSettingsPlacementType(data.GetCurrentPlacementPlacementType())
	}
	return result
}

func PlacementCellClicked(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	chosenCoord := parsedData.Get("chosenCoord")
	// Checking if next or previous type was chosen
	if ships.GetFirstCoord().Coord == "" {
		ships.SetFirstCoord(chosenCoord)
	} else {
		ships.SetLastCoord(chosenCoord)
	}
}

func SwitchPlacementType(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Checking if next or previous type was chosen
	chosenOption := parsedData.Get("chosenOption")
	if chosenOption == "next" {
		SwitchCurrentPlacementType(true)
	} else {
		SwitchCurrentPlacementType(false)
	}
}

func ShipToPlaceChosen(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	chosenOption, err := strconv.Atoi(parsedData.Get("chosenOption"))
	ships.SetPlacingShip(chosenOption)
}
