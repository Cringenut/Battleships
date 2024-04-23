package web

import (
	"fmt"
	"net/http"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}
