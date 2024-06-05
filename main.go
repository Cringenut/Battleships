package main

import (
	"Battleships/data"
	"Battleships/server"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {

	data.InitializePlayerData()
	data.InitializeGameData()
	data.InitializeGameStatus()

	// Initializing gin

	// Using router to handle responses
	// Using New() so we don't get all responses written inside of console
	router := gin.New()
	router.LoadHTMLGlob("views/*")
	html := template.Must(template.ParseFiles("views/battle-page-redirect.html"))
	router.SetHTMLTemplate(html)
	// Gin recovery so server doesn't crash on INTERNAL_ERROR
	router.Use(gin.Recovery())

	// Router for handlers
	app := server.Config{Router: router}
	app.Routes()

	// Load HTML templates
	router.LoadHTMLGlob("views/*")

	router.Run(":8080")

}
