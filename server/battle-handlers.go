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
	"time"
)

func (app *Config) HandleGameStatus(c *gin.Context) {
	if data.GetToken() == "" {
		return
	}

	time.Sleep(100 * time.Millisecond)
	println("status")

	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		Render(c, 200, views.MakeTurnText(false))
		return
	}
	data.SetGameStatus(gameStatus)

	Render(c, 200, views.MakeTurnText(data.IsTurnChanged()))
	data.IsPlayerTurn = gameStatus.ShouldFire
}

func (app *Config) HandleEnemyTarget(c *gin.Context) {
	data.SetEnemyAccuracy(data.CalculateEnemyAccuracy())
	time.Sleep(200 * time.Millisecond)
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandlePlayerTarget(c *gin.Context) {
	if data.GetToken() == "" {
		time.Sleep(200 * time.Millisecond)
		Render(c, 200, views.MakePlayerBoard())
	}
Status:
	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		time.Sleep(500 * time.Millisecond)
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
	data.SetEnemyShots(data.GetGameStatus().OppShots)
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandlePlayerInfo(c *gin.Context) {
	Render(c, 200, views.MakePlayerInfo(data.GetPlayerNickname(), data.GetPlayerDescription()))
}

func (app *Config) HandleEnemyInfo(c *gin.Context) {
	if data.GetToken() == "" {
		return
	}

	enemyData := data.GetEnemyData()
	Render(c, 200, views.MakePlayerInfo(enemyData.Nickname, enemyData.Description))
}

func (app *Config) HandleBattleTimer(c *gin.Context) {
	Render(c, 200, views.MakeBattleTimer(strconv.Itoa(data.GetGameStatus().Timer)))
}

func (app *Config) HandleBattleEnded(c *gin.Context) {
	println("Battle Ended")
	if data.GetGameStatus().LastGameStatus == "win" {
		Render(c, 200, views.MakeWinScreen())
	} else if data.GetGameStatus().LastGameStatus == "lose" {
		Render(c, 200, views.MakeLoseScreen())
	}
}

func (app *Config) HandleEnemyAccuracy(c *gin.Context) {
	Render(c, 200, views.MakeAccuracyField(5.0))
}
