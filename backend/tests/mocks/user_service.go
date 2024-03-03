package mocks

import (
	"go-wallet-api/models"

	"github.com/gofrs/uuid"
)

type MockUserService struct {
	Repository *MockUserRepository
}

// Init new service
func NewMockUserService(repo MockUserRepository) *MockUserService {
	return &MockUserService{Repository: &repo}
}

// This gets all users
func (s *MockUserService) FindAll() ([]models.User, error) {
	return s.Repository.FindAll()
}

func (s *MockUserService) FindById(id uuid.UUID) (models.User, error) {
	return s.Repository.FindById(id)
}

func (s *MockUserService) UserExistsByEmail(email string) bool {
	users, _ := s.FindAll()

	for _, u := range users {
		if u.Email == email {
			return true
		}
	}

	return false
}

func (s *MockUserService) UserExistsByPhone(phone string) bool {
	users, _ := s.FindAll()

	for _, u := range users {
		if u.PhoneNum.String == phone {
			return true
		}
	}

	return false
}
