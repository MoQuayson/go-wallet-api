package database

import (
	"go-wallet-api/features/shared/utils/enums"
	userEntity "go-wallet-api/features/users/business_logic/app/entities"
	walletEntity "go-wallet-api/features/wallet/business_logic/app/entities"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// MigrateEntities Migrate Entities Here
func MigrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.UserEntity{},
		&walletEntity.WalletEntity{},
	)
}

// Seeds admin into db
func SeedUserAdmin(db *gorm.DB) error {
	return db.Save(&userEntity.UserEntity{
		ID:        uuid.FromStringOrNil("81423bfa-3ead-430b-a4f0-de4286ee9dcd"),
		Name:      "Wallet Admin",
		Email:     "admin@example.com",
		Password:  "password",
		Role:      enums.RoleType_Admin,
		CreatedAt: time.Now(),
	}).Error
}
