package tests

import (
	"Battleships/web/ships"
	"testing"
	"time"
)

func TestGenerateRandomCoordinates(t *testing.T) {
	const testIterations = 10
	const timeout = 10 * time.Second

	for i := 0; i < testIterations; i++ {
		done := make(chan bool)
		go func() {
			ships.GenerateRandomCoordinates()
			done <- true
		}()

		select {
		case <-done:
			// Function completed successfully
		case <-time.After(timeout):
			t.Fatalf("Test failed: GenerateRandomCoordinates entered an infinite loop on iteration %d", i)
		}
	}
}
