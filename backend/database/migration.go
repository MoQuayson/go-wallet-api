package database

import (
	"go-wallet-api/models"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Migrate Entity Here
func MigrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
	)
}

// Seeds admin into db
func SeedUserAdmin(db *gorm.DB) error {
	return db.Save(&models.User{
		ID:        uuid.FromStringOrNil("81423bfa-3ead-430b-a4f0-de4286ee9dcd"),
		Name:      "Wallet Admin",
		Email:     "admin@example.com",
		Password:  "password",
		Role:      "Administrator",
		CreatedAt: time.Now(),
	}).Error
}
