package server

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	//views
	app.Router.GET("/", HandleHomePage())

	app.Router.POST("/fire", app.HandleFire)
}
