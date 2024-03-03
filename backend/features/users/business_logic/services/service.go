package services

import (
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/users/business_logic/app/entities"
	"go-wallet-api/features/users/business_logic/app/models"
	"go-wallet-api/features/users/pkg"
)

type UserService struct {
	pkg.IUserService
	repo pkg.IUserRepository
}

func NewUserService(repo pkg.IUserRepository) pkg.IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAllUsers() ([]*models.User, error) {
	usersChan, errChan := utils.MakeDataSliceAndErrorChannels[entities.UserEntity]()
	s.repo.FindAllUsers(usersChan, errChan)
	userEntities := <-usersChan
	err := <-errChan

	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0)
	for _, user := range userEntities {
		users = append(users, models.NewUserModelWithUserEntity(user))
	}

	return users, nil
}
func (s *UserService) FindUserById(id string) (*models.User, error) {
	userChan, errChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	s.repo.FindUserById(id, userChan, errChan)
	userEntity := <-userChan
	err := <-errChan

	if err != nil {
		return nil, err
	}
	return models.NewUserModelWithUserEntity(userEntity), nil
}
func (s *UserService) CreateNewUser(req *models.UserRequest) (*models.User, error) {
	errChan := make(chan error)
	user := models.NewUserEntity(req)
	s.repo.CreateNewUser(user, errChan)
	err := <-errChan

	if err != nil {
		return nil, err
	}
	return models.NewUserModelWithUserEntity(user), nil
}
func (s *UserService) UpdateUser(id string, req *models.UserRequest) (*models.User, error) {
	//get user by id
	userChan, errChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	s.repo.FindUserById(id, userChan, errChan)
	userEntity := <-userChan
	err := <-errChan
	if err != nil {
		return nil, err
	}

	userEntity.Name = req.Name
	userEntity.PhoneNum = &req.PhoneNum
	userEntity.Email = req.Email

	errChan = make(chan error)
	s.repo.UpdateUser(userEntity, errChan)
	if err = <-errChan; err != nil {
		return nil, err
	}

	return models.NewUserModelWithUserEntity(userEntity), nil

}
func (s *UserService) DeleteUser(id string) error {
	errChan := make(chan error)
	s.repo.DeleteUser(id, errChan)
	if err := <-errChan; err != nil {
		return err
	}

	return nil

}
func (s *UserService) CheckIfUserExistByPhone(phoneNum string) (*models.User, error) {
	userChan, errChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	s.repo.CheckIfUserExistByPhone(phoneNum, userChan, errChan)
	userEntity := <-userChan
	err := <-errChan

	if err != nil {
		return nil, err
	}
	return models.NewUserModelWithUserEntity(userEntity), nil
}
