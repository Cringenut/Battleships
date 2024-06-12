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

	// Main menu handlers
	app.Router.POST("/", app.HandleMainMenuContainer)
	app.Router.GET("/battle/redirect", app.HandleMenuRedirectToBattle)
	app.Router.POST("/multiplayer/wait", app.HandleMultiplayerStartWait)
	app.Router.POST("/multiplayer/wait/refresh", app.HandleMultiplayerRefresh)
	app.Router.POST("/multiplayer/lobbies", app.HandleMultiplayerLobbies)
	app.Router.POST("/multiplayer/join", app.HandleMultiplayerJoinLobby)
	app.Router.GET("/multiplayer/wait/check", app.HandleMultiplayerWait)

	// Settings handlers
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

	// Battle handlers
	app.Router.GET("/battle/status", app.HandleGameStatus)
	app.Router.GET("/battle/board/target/enemy", app.HandleEnemyTarget)
	app.Router.GET("/battle/board/target/player", app.HandlePlayerTarget)
	app.Router.POST("/battle/fire", app.HandlePlayerShot)
	app.Router.GET("/battle/shots", app.HandleSetEnemyShots)
	app.Router.POST("/battle/player/info", app.HandlePlayerInfo)
	app.Router.POST("/battle/enemy/info", app.HandleEnemyInfo)
	app.Router.POST("/battle/timer", app.HandleBattleTimer)
	app.Router.POST("/battle/ended", app.HandleBattleEnded)
	app.Router.POST("/battle/enemy/accuracy", app.HandleEnemyAccuracy)
	app.Router.POST("/battle/player/accuracy", app.HandlePlayerAccuracy)
	app.Router.POST("/battle/shots/history", app.HandleShotsHistory)
	app.Router.POST("/battle/surrender/show", app.HandleSurrenderWindowShow)
	app.Router.POST("/battle/surrender", app.HandleSurrender)

	// Errors handlers
	app.Router.GET("/errors/menu", app.HandleMainMenuErrors)
	app.Router.GET("/errors/settings", app.HandleSettingsErrors)
	app.Router.GET("/errors/battle", app.HandleBattleErrors)

}
