package web

import (
	"Battleships/client/data"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	Token string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/battle":
		battlePageHandler(w, r)
	default:

		fmt.Fprint(w, "Default")
	}
}

func battlePageHandler(w http.ResponseWriter, r *http.Request) {
	var fileName = "assets/battle_page.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	pageData := PageData{
		Token: "Game Id: " + data.GetToken(),
	}

	err = t.ExecuteTemplate(w, "battle_page", pageData)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

	paintBoardSquares()

}

func drawBoard() {

}

func paintBoardSquares() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Listen to console events
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *runtime.EventConsoleAPICalled:
			fmt.Println("Console call:")
			for _, arg := range ev.Args {
				if arg.Type == "string" {
					s, err := strconv.Unquote(strings.Replace(strconv.Quote(string(arg.Value)), `\\u`, `\u`, -1))
					if err != nil {
						fmt.Println("Error unquoting string:", err)
					} else {
						fmt.Println("Arg:", s)
					}
				}
			}
		}
	})

	// Run tasks
	var result interface{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://localhost:8080/assets/battle_page.html`),
		chromedp.WaitReady(`#board`, chromedp.ByID),
		chromedp.Evaluate(`paintSquares(["A10","B2","J3"]);`, &result),
	)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Result of paintSquares: %v", result)
	}
}
