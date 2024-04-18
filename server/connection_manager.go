package server

import "net/http"

func CheckConnection() string {
	response, err := http.Get("https://go-pjatk-server.fly.dev/swagger/index.html")
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	return response.Status
}
