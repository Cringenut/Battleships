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

func (app *Config) HandleSave(c *gin.Context) {
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
		return
	} else {
		data.SetPlayerData(saveNickname, saveDescription, web.GetAllShipCoords())
	}

	fmt.Println("Save data:", formData) // Add this line for debugging

	// Respond with an HTML page containing JavaScript to redirect
	c.Header("Content-Type", "text/html")
	c.String(http.StatusMovedPermanently, `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>Redirecting...</title>
        </head>
        <body>
            <p>Processing complete. Redirecting...</p>
            <script type="text/javascript">
                window.location.href = "/";
            </script>
        </body>
        </html>
    `)
	c.Abort() // End the request early
}

func (app *Config) HandlePlacementCell(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Checking if next or previous type was chosen
	chosenCoord := parsedData.Get("chosenCoord")
	if web.GetFirstCoord().Coord == "" {
		web.SetFirstCoord(chosenCoord)
	} else {
		web.SetLastCoord(chosenCoord)
	}

	Render(c, 200, views.MakePlacingElement())
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

	Render(c, 200, views.MakeSettingsPage())
}

func (app *Config) HandleShipPlacementChosen(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Checking if next or previous type was chosen
	chosenOption, err := strconv.Atoi(parsedData.Get("chosenOption"))
	web.SetPlacingShip(chosenOption)
	println(chosenOption)

	Render(c, 200, views.MakePlacingElement())
}
