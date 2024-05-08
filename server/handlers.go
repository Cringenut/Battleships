package server

import (
	"Battleships/client"
	"Battleships/views"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func HandleHomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		render(c, 200, views.MakeBattlePage(client.GetToken()))
	}
}

func (app *Config) HandleFire(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}

	coord := strings.TrimPrefix(string(jsonData), "coord=")
	err = PostFire(coord)
	if err != nil {
		return
	}

	fmt.Println(coord)
	fmt.Println("Bang")

	render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandleGetGameStatus(c *gin.Context) {
	status, _ := GetGameStatus()
	render(c, 200, views.MakeGameStatusFooter(status.ShouldFire, status.Timer))
}

func (app *Config) HandleBoard(c *gin.Context) {
	render(c, 200, views.MakeEnemyBoard())
}
