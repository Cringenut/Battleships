package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// AbandonGame sends a DELETE request to abandon the game
func GameAbandon(token string) error {
	deleteURL := "https://go-pjatk-server.fly.dev/api/game/abandon"

	if err := CheckConnection(); err != nil {
		return err
	}

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to abandon the game")
	}

	return nil
}
