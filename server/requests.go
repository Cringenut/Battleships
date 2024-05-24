package server

import (
	"Battleships/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func CheckConnection() error {
	response, err := http.Get("https://go-pjatk-server.fly.dev/swagger/index.html")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Connection not available")
	}

	return nil
}

/* POST */
func PostInitGame(body []byte) (string, error) {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Handle error
		}
	}(res.Body)

	return res.Header.Get("x-auth-token"), nil
}

/* GET */
func GetBoard(token string) ([]string, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/game/board"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game board")
	}

	var boardResponse data.BoardResponse
	if err := json.NewDecoder(res.Body).Decode(&boardResponse); err != nil {
		return nil, err
	}

	return boardResponse.Board, nil
}
