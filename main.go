package main

import (
	"Battleships/data"
	"Battleships/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println(data.GetPlayerNickname())

	// Initializing gin

	// Using router to handle responses
	// Using New() so we don't get all responses written inside of console
	router := gin.New()
	// Gin recovery so server doesn't crash on INTERNAL_ERROR
	router.Use(gin.Recovery())

	// Router for handlers
	app := server.Config{Router: router}
	app.Routes()

	router.Run(":8080")

}
