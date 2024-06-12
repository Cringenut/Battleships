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
	app.Router.GET("/ranking", app.HandleRankingPage)

	// Handlers for main menu
	app.Router.POST("/", app.HandleMainMenuContainer)
	app.Router.POST("/multiplayer/wait", app.HandleMultiplayerStartWait)
	app.Router.POST("/multiplayer/wait/refresh", app.HandleMultiplayerRefresh)
	app.Router.POST("/multiplayer/lobbies", app.HandleMultiplayerLobbies)
	app.Router.POST("/multiplayer/join", app.HandleMultiplayerJoinLobby)
	app.Router.GET("/battle/redirect", app.HandleMenuRedirectToBattle)
	app.Router.GET("/multiplayer/wait/check", app.HandleMultiplayerWait)

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
	app.Router.POST("/placement/show", app.HandlePlacementShow)
	app.Router.GET("/placement/randomise", app.HandlePlacementRandomise)
	app.Router.POST("/placement/back")

	// Handlers for battle
	app.Router.GET("/status", app.HandleGameStatus)
	app.Router.GET("/board/target/enemy", app.HandleEnemyTarget)
	app.Router.GET("/board/target/player", app.HandlePlayerTarget)
	app.Router.POST("/handle/fire", app.HandlePlayerShot)
	app.Router.GET("/shots", app.HandleSetEnemyShots)
	app.Router.POST("/player/info", app.HandlePlayerInfo)
	app.Router.POST("/enemy/info", app.HandleEnemyInfo)
	app.Router.POST("/timer", app.HandleBattleTimer)
	app.Router.POST("/battle/ended", app.HandleBattleEnded)
	app.Router.POST("/battle/enemy/accuracy", app.HandleEnemyAccuracy)
	app.Router.POST("/battle/player/accuracy", app.HandlePlayerAccuracy)
	app.Router.POST("/battle/shots/history", app.HandleShotsHistory)

}
