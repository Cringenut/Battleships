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
	app.Router.POST("/multiplayer/wait", app.HandleMultiplayerWait)
	app.Router.POST("/multiplayer/wait/refresh", app.HandleMultiplayerRefresh)
	app.Router.POST("/multiplayer/lobbies", app.HandleMultiplayerLobbies)

	// Handlers for settings
	app.Router.POST("/settings/save", app.HandleSettingsSave)
	app.Router.POST("/placement/page")
	app.Router.POST("/placement/place", app.HandlePlacementCell)
	app.Router.POST("/placement/advanced/place", app.HandlePlacementCell)
	app.Router.POST("/placement/switch", app.HandlePlacementTypeSwitch)
	app.Router.POST("/placement/option", app.HandlePlacementChosen)
	app.Router.POST("/placement/clear", app.HandlePlacementClear)
	app.Router.POST("/placement/cancel", app.HandlePlacementCancel)
	app.Router.POST("/placement/save", app.HandlePlacementSave)
	app.Router.POST("/placement/back", app.HandlePlacementBack)
	app.Router.POST("/placement/show", app.HandlePlacementShow)

	// Handlers for battle
	app.Router.GET("/status", app.HandleGameStatus)
	app.Router.GET("/board/turn/player", app.HandlePlayerTurn)
	app.Router.GET("/board/turn/enemy", app.HandleEnemyTurn)
	app.Router.POST("/handle/fire", app.HandleFire)
	app.Router.GET("/shots", app.HandleSetShots)
	app.Router.POST("/player/info", app.HandlePlayerInfo)
	app.Router.POST("/enemy/info", app.HandleEnemyInfo)
	app.Router.POST("/timer", app.HandleBattleTimer)

}
