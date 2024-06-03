package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (app *Config) HandleSettingsSave(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to read request body")
		return
	}

	// Parse the form data
	formData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to parse form data")
		return
	}

	saveNickname := formData.Get("nickname")
	saveDescription := formData.Get("description")

	fmt.Println("NICKNAME: " + saveNickname)
	fmt.Println("DESCRIPTION: " + saveDescription)

	if saveNickname == "" {
		Render(c, 200, views.MakeSettingsPage("", saveDescription))
		return
	} else {
		data.SetPlayerData(saveNickname, saveDescription)
		data.SetPlayerShipPlacementType(web.GetCurrentSettingsPlacementType())
	}

	// Respond with an HTML page containing HTML and javascript to redirect
	c.Header("Content-Type", "text/html")
	//c.Redirect(http.StatusTemporaryRedirect, "/")
	c.HTML(http.StatusOK, "redirect.html", gin.H{})
	c.Abort()
}

func (app *Config) HandlePlacementCell(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	chosenCoord := parsedData.Get("chosenCoord")
	// Checking if next or previous type was chosen
	if web.GetFirstCoord().Coord == "" {
		web.SetFirstCoord(chosenCoord)
	} else if web.GetCurrentPlacementPlacementType() == data.Advanced && len(web.GetEndCoords()) != 0 && len(web.GetPlacingShip().Coords) < web.GetPlacingShip().Size-1 && web.GetPlacingShip().Size > 2 {
		web.SetNextCoord(chosenCoord)
	} else {
		web.SetLastCoord(chosenCoord)
	}

	Render(c, 200, views.MakePlacementElement())
}

func (app *Config) HandlePlacementTypeSwitch(c *gin.Context) {
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
		web.SwitchCurrentPlacementType(true)
	} else {
		web.SwitchCurrentPlacementType(false)
	}

	web.SetPlacingShip(-1)
	Render(c, 200, views.MakeShipPlacementElement())
}

func (app *Config) HandlePlacementChosen(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Checking if next or previous type was chosen
	chosenOption, err := strconv.Atoi(parsedData.Get("chosenOption"))
	isAdvanced, err := strconv.ParseBool(parsedData.Get("isAdvanced"))
	web.SetPlacingShip(chosenOption)
	println(isAdvanced)

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
	web.SetCurrentSettingsPlacementType(web.GetCurrentPlacementPlacementType())
	println(data.GetPlayerShipPlacementType())
}

func (app *Config) HandlePlacementBack(c *gin.Context) {
	return
}

func (app *Config) HandlePlacementShow(c *gin.Context) {
	web.SetCurrentPlacementPlacementType(web.GetCurrentSettingsPlacementType())
	web.SetPlacingShip(-1)
	Render(c, 200, views.MakeShipPlacementElement())
}
