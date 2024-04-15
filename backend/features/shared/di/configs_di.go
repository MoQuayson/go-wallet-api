package di

import "go-wallet-api/configs"

var WithConfigsInjector *ConfigsInjector

type ConfigsInjector struct {
	KeyStoreConfig *configs.KeyStoreConfig
}

func NewConfigsInjector() *ConfigsInjector {
	injector := &ConfigsInjector{}
	injector.KeyStoreConfig = configs.NewKeyStoreConfig()

	return injector
}
