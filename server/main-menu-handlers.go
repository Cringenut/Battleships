package server

import (
	"Battleships/data"
	"Battleships/requests"
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
		// If battle didn't start we try to start it again until it returns no error
	Battle:
		err := web.StartBattle("", true)
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
	web.StartBattle("", false)
	Render(c, 200, views.MakeMultiplayerWaitChosen())
}

func (app *Config) HandleMultiplayerRefresh(c *gin.Context) {
	println("Start refresh")
	err := requests.GetGameRefresh(data.GetToken())
	if err != nil {
		return
	}
	println("Refreshed")
}

func (app *Config) HandleMultiplayerLobbies(c *gin.Context) {
	servers, err := requests.GetLobby()
	if err != nil {
		return
	}

	for index, server := range servers {
		Render(c, 200, views.MakePlayerLobby(server.Nick, index))
	}

	println(len(servers))
	println("Lobbies")
}

func (app *Config) HandleMultiplayerJoinLobby(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Checking if next or previous type was chosen
	chosenLobby := parsedData.Get("chosenLobby")
	println("Chosen lobby: " + chosenLobby)

	err = web.StartBattle(chosenLobby, false)
	app.HandleMenuRedirectToBattle(c)

	if err != nil {
		println("ERROR: " + err.Error())
	}
}

func (app *Config) HandleMenuRedirectToBattle(c *gin.Context) {
	println("Redirect to battle")
	web.RedirectToBattle(c)
}

func (app *Config) HandleMultiplayerWait(c *gin.Context) {
	println("Checking battle")
	status, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		return
	}

	if status.Opponent != "" {
		app.HandleMenuRedirectToBattle(c)
	}
}
