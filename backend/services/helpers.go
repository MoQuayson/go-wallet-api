package services

import (
	"go-wallet-api/models"
	"os"
	"sync"
)

// This function waits till all channel ops are completed before closing the channel
func WaitAndCloseChannel(wg *sync.WaitGroup, channel chan models.DBResponse) {
	go func(wg *sync.WaitGroup, channel chan models.DBResponse) {
		wg.Wait()
		close(channel)
	}(wg, channel)
}

// Gets JWT Secret from env
func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
