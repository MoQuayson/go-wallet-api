package utils

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

func NewUUID() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}

func MakeDataAndErrorChannels[T any]() (chan *T, chan error) {
	return make(chan *T, 1), make(chan error, 1)
}

func MakeDataSliceAndErrorChannels[T any]() (chan []*T, chan error) {
	return make(chan []*T, 1), make(chan error, 1)
}

func GetDataFromChannel[T any](dataChan chan *T) *T {
	return <-dataChan
}

func GenerateWalletName(accountScheme, accountType string) string {
	return fmt.Sprintf("%s %s wallet", accountScheme, strings.ToUpper(accountType))
}
