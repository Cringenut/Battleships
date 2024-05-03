package client

import "sync"

var (
	tokenMutex  sync.RWMutex // to ensure safe concurrent access to the token
	token       string
	playerShips []string
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
func SetShips(s []string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	playerShips = s
}

// GetToken returns the token.
func GetShips() []string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return playerShips
}
