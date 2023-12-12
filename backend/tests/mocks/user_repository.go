package mocks

import (
	"go-wallet-api/models"

	"github.com/gofrs/uuid"
)

type MockUserRepository struct{}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	mockUsers := []models.User{
		{ID: models.GenerateUUID(), Name: "Jane Doe", Email: "jane.doe@example.com", PhoneNum: models.ConvertToNullString("069948828929882"), Role: "User"},
		{ID: models.GenerateUUID(), Name: "James Wilson", Email: "james.thewilson19875@example.com", PhoneNum: models.ConvertToNullString("006696884434"), Role: "User"},
		{ID: models.GenerateUUID(), Name: "John Doe", Email: "john.doe@example.com", PhoneNum: models.ConvertToNullString("2335548775889"), Role: "Admin"},
		{ID: models.GenerateUUID(), Name: "Kofi Manu", Email: "manu.kofi@example.com", Role: "User"},
		{ID: models.GenerateUUID(), Name: "Tony Stark", Email: "tony.stark@example.com", Role: "Admin"},
	}

	return mockUsers, nil
}
func (m *MockUserRepository) FindById(id uuid.UUID) (models.User, error) {
	return models.User{
		ID:       id,
		Name:     "Jane Doe",
		Email:    "jane.doe@example.com",
		PhoneNum: models.ConvertToNullString("069948828929882"),
		Role:     "User",
	}, nil
}
