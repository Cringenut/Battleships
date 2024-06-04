package server

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/views"
	"Battleships/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
		// Respond with an HTML page containing HTML and javascript to redirect
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusTemporaryRedirect, "battle-page-redirect.html", gin.H{})
		c.Abort()
	case "multiplayer":
		Render(c, 200, views.MakeLobbiesList())
	case "back":
		Render(c, 200, views.MakeMainMenu())
	default:
		Render(c, 200, views.MakeMainMenu())
	}

	fmt.Println(chosenOption)
}

func (app *Config) HandleMultiplayerWait(c *gin.Context) {
	web.MultiplayerWaitForOpponent()
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
