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
	data.SetToken(res.Header.Get("x-auth-token"))

	if ships, err := GetBoard(); err != nil {
		return err
	} else {
		data.SetPlayerShips(ships)
	}
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
	r.Header.Add("X-Auth-Token", data.GetToken()) // Assuming authToken is provided when function is called

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	responseData := &data.FireResponseData{}
	err = json.NewDecoder(res.Body).Decode(responseData)
	if err != nil {
		return err
	}

	// Print or use the "result" parameter
	fmt.Println("Response from server:", responseData.Result)

	if responseData.Result == "" {
		return nil
	}

	if responseData.Result == "miss" {
		data.AppendPlayerShots(coord, false)
	} else {
		data.AppendPlayerShots(coord, true)
	}

	return nil
}

/* GET */
func GetBoard() ([]string, error) {
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
	req.Header.Add("X-Auth-Token", data.GetToken())

	// Create a new HTTP data and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game board")
	}

	// Decode the JSON response into the GetBoardResponseData struct
	responseData := &data.GetBoardResponseData{}
	err = json.NewDecoder(res.Body).Decode(responseData)
	if err != nil {
		return nil, err
	}

	// Return the board data
	return responseData.Board, nil
}

func GetGameStatus() (*data.GetGameStatusData, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", data.GetToken())

	// Create a new HTTP data and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game board")
	}

	// Decode the JSON response into the GetBoardResponseData struct
	responseData := &data.GetGameStatusData{}
	err = json.NewDecoder(res.Body).Decode(responseData)
	if err != nil {
		return nil, err
	}

	// Return the board data
	return responseData, nil
}
