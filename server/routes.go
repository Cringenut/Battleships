package server

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	// Main page getters
	app.Router.GET("/", app.HandleMainMenu)
	app.Router.GET("/battle", app.HandleBattlePage)
	app.Router.GET("/settings", app.HandleSettings)

	// Handlers for main menu
	app.Router.POST("/", app.HandleMainMenuContainer)
	app.Router.GET("/redirect", app.HandleBattlePageRedirect)

	// Handlers for settings
	app.Router.POST("/save", app.HandleSave)
	app.Router.POST("/place", app.HandlePlacementCell)

	// Handlers for battle
	app.Router.GET("/status", app.HandleGameStatus)
	app.Router.GET("/board/turn/player", app.HandlePlayerTurn)
	app.Router.GET("/board/turn/enemy", app.HandleEnemyTurn)
	app.Router.POST("/handle/fire", app.HandleFire)
	app.Router.GET("/shots", app.HandleSetShots)
	app.Router.POST("/player/info", app.HandlePlayerInfo)
	app.Router.POST("/enemy/info", app.HandleEnemyInfo)

}
