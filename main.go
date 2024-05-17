package main

import (
	"Battleships/server"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	//initialize config
	app := server.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8080")

}
