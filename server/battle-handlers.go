package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web/battle"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (app *Config) HandleGameStatus(c *gin.Context) {
	battle.CheckGameStatus()
	Render(c, 200, views.MakeTurnText())
}

func (app *Config) HandleEnemyTarget(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandlePlayerTarget(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	Render(c, 200, views.MakePlayerBoard())
}

func (app *Config) HandlePlayerShot(c *gin.Context) {
	battle.FireAtEnemy(c)
	battle.CalculatePlayerAccuracy()
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandleSetEnemyShots(c *gin.Context) {
	data.SetEnemyShots(data.GetGameStatus().OppShots)
	battle.CalculateEnemyAccuracy()
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
	if battle.CheckWin() {
		Render(c, 200, views.MakeWinScreen())
	} else {
		Render(c, 200, views.MakeLoseScreen())
	}
}

func (app *Config) HandleEnemyAccuracy(c *gin.Context) {
	Render(c, 200, views.MakeEnemyAccuracyField())
}

func (app *Config) HandlePlayerAccuracy(c *gin.Context) {
	Render(c, 200, views.MakePlayerAccuracyField())
}
