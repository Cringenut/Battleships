package server

import (
	"Battleships/client"
	"Battleships/views"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
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

func HandleFire() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Handled")
	}
}
