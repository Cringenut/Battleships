package server

import (
	"Battleships/views"
	"Battleships/web"
	"Battleships/web/battle"
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
		// If battle didn't start we try to start it again until it returns no error
	Battle:
		err := battle.StartBattle("", true)
		if err != nil {
			goto Battle
		}
		Render(c, 200, views.MakeSingleplayerChosen())
	case "multiplayer":
		Render(c, 200, views.MakeLobbiesList())
	case "back":
		Render(c, 200, views.MakeMainMenu())
	default:
		Render(c, 200, views.MakeMainMenu())
	}

	fmt.Println(chosenOption)
}

func (app *Config) HandleMultiplayerStartWait(c *gin.Context) {
	battle.StartBattle("", false)
	Render(c, 200, views.MakeMultiplayerWaitChosen())
}

func (app *Config) HandleMultiplayerRefresh(c *gin.Context) {
	web.RefreshLobby()
}

func (app *Config) HandleMultiplayerLobbies(c *gin.Context) {
	for index, server := range web.FindLobbies() {
		Render(c, 200, views.MakePlayerLobby(server.Nick, index))
	}
}

func (app *Config) HandleMultiplayerJoinLobby(c *gin.Context) {
	web.JoinPlayerLobby(c)
}

func (app *Config) HandleMenuRedirectToBattle(c *gin.Context) {
	println("Redirect to battle")
	web.RedirectToBattle(c)
}

func (app *Config) HandleMultiplayerWait(c *gin.Context) {
	println("Checking battle")
	web.CheckIfSomeoneJoinedLobby(c)
}
