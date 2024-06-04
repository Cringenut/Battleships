package server

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/views"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"strings"
)

func (app *Config) HandleGameStatus(c *gin.Context) {
	if data.GetToken() == "" {
		return
	}

	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		Render(c, 200, views.MakeTurnText(false))
		return
	}
	data.SetGameStatus(gameStatus)

	println("Time: " + strconv.Itoa(gameStatus.Timer))

	Render(c, 200, views.MakeTurnText(data.IsTurnChanged()))
	data.IsPlayerTurn = gameStatus.ShouldFire
}

func (app *Config) HandlePlayerTurn(c *gin.Context) {
	println("Player")
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandleEnemyTurn(c *gin.Context) {
Status:
	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		goto Status
	}
	data.SetGameStatus(gameStatus)
	Render(c, 200, views.MakePlayerBoard())
}

func (app *Config) HandleFire(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Render(c, 200, views.MakeEnemyBoard())
	}

	coord := strings.TrimPrefix(string(jsonData), "coord=")
	res, err := requests.PostFire(data.GetToken(), coord)
	if err != nil {
		Render(c, 200, views.MakeEnemyBoard())
		return
	}

	fmt.Println(coord)
	data.AppendPlayerShots(coord, res)

	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandleSetShots(c *gin.Context) {
	data.AppendEnemyShotsToHistory()
	data.SetEnemyShots(data.GetGameStatus().OppShots)
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandlePlayerInfo(c *gin.Context) {
	println("PlayerInfo")
	Render(c, 200, views.MakePlayerInfo(data.GetPlayerNickname(), data.GetPlayerDescription()))
}

func (app *Config) HandleEnemyInfo(c *gin.Context) {
	println("EnemyInfo")
	if data.GetToken() == "" {
		return
	}

	enemyData := data.GetEnemyData()
	Render(c, 200, views.MakePlayerInfo(enemyData.Nickname, enemyData.Description))
}

func (app *Config) HandleBattleTimer(c *gin.Context) {

}
