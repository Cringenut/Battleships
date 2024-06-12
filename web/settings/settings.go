package settings

import (
	"Battleships/data"
	"Battleships/web/errors"
	"Battleships/web/redirect"
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

	// Getting values from the html file
	saveNickname := formData.Get("nickname")
	saveDescription := formData.Get("description")

	// Debug
	fmt.Println("NICKNAME: " + saveNickname)
	fmt.Println("DESCRIPTION: " + saveDescription)

	// If player nickname is empty stay on the page
	if saveNickname == "" {
		errors.AddSettingsError("Nickname is required")
		return saveDescription
	}

	data.SetPlayerData(saveNickname, saveDescription)
	if !ships.IsAnyShipMissingCoords() {
		data.SetPlayerShipPlacementType(data.GetCurrentSettingsPlacementType())
		// The GetAllShipCoords() uses the current placement of placement element not settings
		// In order to not create additional function setting the placement for the placement element ourselves
		// If we use open the placement element there will be no issues, the placement type is set on open from settings
		data.SetCurrentPlacementPlacementType(data.GetCurrentSettingsPlacementType())

		data.SetPlayerShips(ships.GetAllShipsCoords())
		redirect.Redirect(c, "/")
	} else {
		errors.AddSettingsError("Current ship placement missing coordinates")
	}

	return saveDescription
}

// Simple carosuel to switch between placement types if arrow button is clicked
func SwitchCurrentPlacementType(isNext bool) {
	currentIndex := findPlacementIndex(data.GetCurrentPlacementPlacementType())
	if currentIndex == -1 {
		fmt.Println("Current placement type not found.")
		errors.AddSettingsError("Current placement type not found")
		return
	}

	// If last index go to 0
	if isNext {
		if currentIndex+1 >= len(data.GetPlacementTypes()) {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[0])
		} else {
			data.SetCurrentPlacementPlacementType(data.GetPlacementTypes()[currentIndex+1])
		}
		// If 0 go to last
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
	case data.Simple:
		// Check if any of the ships doesn't have all their coordinates
		result = !ships.IsAnyShipMissingCoords()
	case data.Advanced:
		result = !ships.IsAnyShipMissingCoords()
	// Next placement types always filled in the right way so no check required
	case data.Random:
		return true
	case data.ServerRandom:
		return true
	default:
		result = false
	}

	// Setting the currentSettingsPlacementType for the settings page
	if result == true {
		data.SetCurrentSettingsPlacementType(data.GetCurrentPlacementPlacementType())
	} else {
		errors.AddSettingsError("Current ship placement missing coordinates")
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
	// If no first coord set it
	if ships.GetFirstCoord().Coord == "" {
		ships.SetFirstCoord(chosenCoord)
		// Used only for advanced placement
		// The length of not filled advanced ship would always fulfil this condition
		// Plus represent adding first and last coord
		// Would be false when last coord should be placed
	} else if data.GetCurrentPlacementPlacementType() == data.Advanced && len(ships.GetNextCoords())+2 < ships.GetPlacingShip().Size {
		ships.SetNextCoord(chosenCoord)
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
	ships.ClearBoard()
	ships.RepopulateBoard()
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
