package client

import "sync"

type PlayerShot struct {
	coord string
	hit   bool
}

var (
	tokenMutex  sync.RWMutex // to ensure safe concurrent access to the token
	token       string
	playerShips []string
	playerShots []PlayerShot
)

// SetToken sets the token.
func SetToken(t string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	token = t
}

// GetToken returns the token.
func GetToken() string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return token
}

// SetToken sets the token.
func SetPlayerShips(s []string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	playerShips = s
}

// GetToken returns the token.
func GetPlayerShips() []string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return playerShips
}

func AppendPlayerShots(s PlayerShot) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	playerShots = append(playerShots, s)
}
