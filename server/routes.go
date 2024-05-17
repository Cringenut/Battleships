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

	// Handlers for main menu
	app.Router.POST("/", app.HandleMainMenuContainer)
	app.Router.GET("/redirect", app.HandleBattlePageRedirect)
}
