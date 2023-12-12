package mocks

import (
	"go-wallet-api/models"

	"github.com/gofrs/uuid"
)

type MockUserService struct {
	Repository *MockUserRepository
}

func NewMockUserService(repo MockUserRepository) *MockUserService {
	return &MockUserService{Repository: &repo}
}

func (s *MockUserService) FindAll() ([]models.User, error) {
	return s.Repository.FindAll()
}

func (s *MockUserService) FindById(id uuid.UUID) (models.User, error) {
	return s.Repository.FindById(id)
}
