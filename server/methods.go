package server

import (
	"Battleships/client"
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
func PostInitGame(body []byte) error {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer((body)))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	// If post goes successful
	client.SetToken(res.Header.Get("x-auth-token"))
	return nil
}

func PostFire(coord string) error {
	posturl := "https://go-pjatk-server.fly.dev/api/game/fire"

	// Creating the JSON body using a map and marshal it to bytes
	jsonData := map[string]string{"coord": coord}
	body, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	if err := CheckConnection(); err != nil {
		return err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Adding the necessary headers
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Auth-Token", client.GetToken()) // Assuming authToken is provided when function is called

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Handle the error
		}
	}(res.Body)

	// Read the response body
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Print the response body
	fmt.Println("Response from server:", string(responseData))

	return nil
}
