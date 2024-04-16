package seeders

import (
	"github.com/gofrs/uuid"
	"go-wallet-api/features/shared/di"
	"go-wallet-api/features/shared/utils/enums"
	userEnt "go-wallet-api/features/users/business_logic/entities"
	"gorm.io/gorm"
	"time"
)

var (
	adminId = uuid.FromStringOrNil("81423bfa-3ead-430b-a4f0-de4286ee9dcd")
)

// SeedUserAdmin Seeds admin data into db
func SeedUserAdmin() error {
	db := di.WithSharedInjector.DB

	errChan := make(chan error, 1)

	go func(*gorm.DB, chan error) {
		err := db.Where("id = ?", adminId).FirstOrCreate(&userEnt.UserEntity{
			ID:        adminId,
			Name:      "Wallet Admin",
			Email:     "admin@example.com",
			Password:  "password",
			Role:      enums.RoleType_Admin,
			CreatedAt: time.Now(),
		}).Error

		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}(db, errChan)

	return <-errChan
}
