package server

import (
	"Battleships/views"
	"Battleships/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
)

func (app *Config) HandleMainMenuContainer(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Extracting the chosen option from the parsed data
	chosenOption := parsedData.Get("chosenOption")

	switch chosenOption {
	case "battle":
		return
	case "single":
		Render(c, 200, views.MakeSingleplayerChosen())
		err := web.StartBattle()
		if err != nil {
			return
		}
	case "back":
		Render(c, 200, views.MakeMainMenu())
	default:
		Render(c, 200, views.MakeMainMenu())
	}

	fmt.Println(chosenOption)
}

func (app *Config) HandleBattlePageRedirect(c *gin.Context) {
	web.CheckBattleDataIntegrity()
	fmt.Println("Redirect")
}