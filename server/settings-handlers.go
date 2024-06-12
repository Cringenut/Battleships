package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web/settings"
	"Battleships/web/ships"
	"github.com/gin-gonic/gin"
)

// Make it so description stays the same after the page is refreshed when settings couldn't be saved
func (app *Config) HandleSettingsSave(c *gin.Context) {
	description := settings.SaveSettings(c)
	Render(c, 200, views.MakeSettingsPage("", description))
}

// Handle click at placement cell
func (app *Config) HandlePlacementCell(c *gin.Context) {
	settings.PlacementCellClicked(c)
	Render(c, 200, views.MakePlacementElement())
}

// Switch placement placement type switch
func (app *Config) HandlePlacementTypeSwitch(c *gin.Context) {
	settings.SwitchPlacementType(c)
	// Clearing all the placement data to avoid visual bugs
	ships.SetPlacingShip(-1)
	Render(c, 200, views.MakeShipPlacementElement())
}

// Choosing ship to place
func (app *Config) HandlePlacementChosen(c *gin.Context) {
	settings.ShipToPlaceChosen(c)
	Render(c, 200, views.MakePlacementElement())
}

// Clearing all placed ships depending on current placement type
func (app *Config) HandlePlacementClear(c *gin.Context) {
	ships.ClearAllShipsCoords()
	Render(c, 200, views.MakePlacementElement())
}

// Clear placing ship
func (app *Config) HandlePlacementCancel(c *gin.Context) {
	ships.SetPlacingShip(-1)
	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementSave(c *gin.Context) {
	// If ships are not missing any coord save the placement
	if !settings.CanCurrentPlacementBeSaved() {
		Render(c, 200, views.MakeShipPlacementElement())
	}
	// If save successful set placement type for settings
	data.SetCurrentSettingsPlacementType(data.GetCurrentPlacementPlacementType())
	println(data.GetPlayerShipPlacementType())
}

// Showing the placement board depending on current placement
func (app *Config) HandlePlacementShow(c *gin.Context) {
	data.SetCurrentPlacementPlacementType(data.GetCurrentSettingsPlacementType())
	Render(c, 200, views.MakeShipPlacementElement())
}

// Randomise ships when randomise button is clicked
func (app *Config) HandlePlacementRandomise(c *gin.Context) {
	ships.GenerateRandomCoordinates()
	Render(c, 200, views.MakePlacementElement())
}
