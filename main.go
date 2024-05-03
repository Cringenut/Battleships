package main

import (
	"Battleships/client"
	"Battleships/game"
	"Battleships/pregame"
	"Battleships/views"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"time"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {

	requestBody := pregame.BuildPostBody()

	if err := game.InitGame(requestBody); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
		fmt.Println("Reconnection...")
	}

	fmt.Println("Game token is: " + client.GetToken())
	_, coords := game.GetBoard()
	client.SetShips(coords)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.MakeBattlePage(client.GetToken()))
	})
	r.Run(":8080")

}
