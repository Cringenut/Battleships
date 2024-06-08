package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web"
	"github.com/gin-gonic/gin"
)

func (app *Config) HandleSettingsSave(c *gin.Context) {
	description := web.SaveSettings(c)
	Render(c, 200, views.MakeSettingsPage("", description))
}

func (app *Config) HandlePlacementCell(c *gin.Context) {
	web.PlacementCellClicked(c)
	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementTypeSwitch(c *gin.Context) {
	web.SwitchPlacementType(c)
	web.SetPlacingShip(-1)
	Render(c, 200, views.MakeShipPlacementElement())
}

func (app *Config) HandlePlacementChosen(c *gin.Context) {
	web.ShipToPlaceChosen(c)
	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementClear(c *gin.Context) {
	web.ClearAllShipCoords()
	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementCancel(c *gin.Context) {
	web.SetPlacingShip(-1)
	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementSave(c *gin.Context) {
	if !web.CanCurrentPlacementBeSaved() {
		Render(c, 200, views.MakeShipPlacementElement())
	}
	data.SetCurrentSettingsPlacementType(data.GetCurrentPlacementPlacementType())
	println(data.GetPlayerShipPlacementType())
}

func (app *Config) HandlePlacementShow(c *gin.Context) {
	data.SetCurrentPlacementPlacementType(data.GetCurrentSettingsPlacementType())
	Render(c, 200, views.MakeShipPlacementElement())
}
