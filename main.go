package main

import (
	"Battleships/data"
	"Battleships/server"
	"github.com/gin-gonic/gin"
)

func main() {

	data.InitializePlayerData()
	data.InitializeCurrentGameData()

	// Initializing gin

	// Using router to handle responses
	// Using New() so we don't get all responses written inside of console
	router := gin.New()
	// Gin recovery so server doesn't crash on INTERNAL_ERROR
	router.Use(gin.Recovery())

	// Router for handlers
	app := server.Config{Router: router}
	app.Routes()

	// Load HTML templates
	router.LoadHTMLGlob("views/*")

	router.Run(":8080")

}
