package di

import (
	"github.com/go-playground/validator/v10"
	"go-wallet-api/features/shared/utils"
	"gorm.io/gorm"
)

var WithSharedInjector *SharedInjector

type SharedInjector struct {
	DB        *gorm.DB
	Validator *validator.Validate
}

func NewSharedInjector(cfg *ConfigsInjector) *SharedInjector {
	injector := &SharedInjector{}
	injector.DB = cfg.KeyStoreConfig.DBContext
	injector.Validator = utils.NewValidation()

	return injector
}
