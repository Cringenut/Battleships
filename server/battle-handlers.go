package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web/battle"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

// Turn text is responsible for calling this handler and updating the game status
func (app *Config) HandleGameStatus(c *gin.Context) {
	battle.CheckGameStatus()
	Render(c, 200, views.MakeTurnText())
}

// Update enemy board
func (app *Config) HandleEnemyTarget(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	println("Enemy board")
	Render(c, 200, views.MakeEnemyBoard())
}

// Update player board
func (app *Config) HandlePlayerTarget(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	Render(c, 200, views.MakePlayerBoard())
}

// Called when clicked on enemy cell
func (app *Config) HandlePlayerShot(c *gin.Context) {
	battle.FireAtEnemy(c)
	// After each shot calculate the player accuracy
	battle.CalculatePlayerAccuracy()
	Render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandleSetEnemyShots(c *gin.Context) {
	battle.CheckEnemyShots()
	// Constantly update enemy accuracy so no shots would be miscalculated
	battle.CalculateEnemyAccuracy()
	Render(c, 200, views.MakeEnemyBoard())
}

// Show player info window during the battle
func (app *Config) HandlePlayerInfo(c *gin.Context) {
	Render(c, 200, views.MakePlayerInfo(data.GetPlayerNickname(), data.GetPlayerDescription()))
}

// Show enemy info window during the battle
func (app *Config) HandleEnemyInfo(c *gin.Context) {
	if data.GetToken() == "" {
		return
	}

	enemyData := data.GetEnemyData()
	Render(c, 200, views.MakePlayerInfo(enemyData.Nickname, enemyData.Description))
}

// Updating the timer using the game status
func (app *Config) HandleBattleTimer(c *gin.Context) {
	Render(c, 200, views.MakeBattleTimer(strconv.Itoa(data.GetGameStatus().Timer)))
}

// If battle ended replace everything with win or lose screen
// So nothing would be triggered
func (app *Config) HandleBattleEnded(c *gin.Context) {
	if battle.CheckWin() {
		Render(c, 200, views.MakeWinScreen())
	} else {
		Render(c, 200, views.MakeLoseScreen())
	}
}

// Create and update enemy accuracy field
func (app *Config) HandleEnemyAccuracy(c *gin.Context) {
	Render(c, 200, views.MakeEnemyAccuracyField())
}

// Create and update player accuracy field
func (app *Config) HandlePlayerAccuracy(c *gin.Context) {
	Render(c, 200, views.MakePlayerAccuracyField())
}

// Add all shots from history to the side menu
func (app *Config) HandleShotsHistory(c *gin.Context) {
	for index, shot := range data.GetShotsHistory() {
		shot = data.GetShotsHistory()[len(data.GetShotsHistory())-1-index]
		Render(c, 200, views.MakeShotsHistoryItem(shot.Shot.Coord, strings.Title(shot.Shot.ShotResult), shot.Owner))
	}
}
