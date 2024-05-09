package server

import (
	"Battleships/data"
	"Battleships/views"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func HandleHomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Homepage")
		render(c, 200, views.MakeBattlePage(data.GetToken()))
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
	gameData, err := GetGameStatus()
	if err != nil {
		data.PrintErrorInfo(err)
		for err != nil {
			time.Sleep(250 * time.Millisecond)
			gameData, err = GetGameStatus()
		}
	}
	data.SetGameData(gameData)
	render(c, 200, views.MakeGameStatusFooter(data.GetGameData().ShouldFire, data.GetGameData().Timer))
}

func (app *Config) HandleEnemyBoard(c *gin.Context) {
	render(c, 200, views.MakeEnemyBoard())
}

func (app *Config) HandlePlayerBoard(c *gin.Context) {
	render(c, 200, views.MakePlayerBoard())
}
