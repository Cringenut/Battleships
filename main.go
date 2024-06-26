package main

import (
	"Battleships/data"
	"Battleships/server"
	"Battleships/web/ships"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initializing default data
	data.InitializePlayerData()
	data.InitializeGameData()
	data.InitializeGameStatus()
	ships.GenerateRandomCoordinates()

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
