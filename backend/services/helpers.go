package services

import (
	"go-wallet-api/models"
	"sync"
)

func WaitAndCloseChannel(wg *sync.WaitGroup, channel chan models.DBResponse) {
	go func(wg *sync.WaitGroup, channel chan models.DBResponse) {
		wg.Wait()
		close(channel)
	}(wg, channel)
}
