package units

import (
	"go-wallet-api/tests/mocks"
	"testing"

	"github.com/gofrs/uuid"
)

func TestUserService_FindAll(t *testing.T) {
	repo := mocks.MockUserRepository{}
	svc := mocks.NewMockUserService(repo)

	users, err := svc.Repository.FindAll()
	expectedUsers := 5

	if err != nil {
		t.Errorf("Expect no error got:%v", err)
	}

	if expectedUsers != len(users) {
		t.Errorf("FAILED:Expected %d users got:%d", expectedUsers, len(users))
	} else {
		t.Logf("Success: Expected %d users got:%d", expectedUsers, len(users))
	}

}
func TestUserService_FindUserById(t *testing.T) {
	repo := mocks.MockUserRepository{}
	svc := mocks.NewMockUserService(repo)

	uid := uuid.FromStringOrNil("538a7239-677e-4ca4-8b22-c6746380c4d6")
	user, err := svc.Repository.FindById(uid)
	expectedUserId := uid

	if err != nil {
		t.Errorf("Expected no error got:%v", err)
	}

	if expectedUserId != user.ID {
		t.Errorf("Expected userId:%v  got:%v", expectedUserId, user.ID)
	} else {
		t.Logf("Success: Expected userId:%v  got:%v", expectedUserId, user.ID)
	}

}

// Test to check if user exist by email
func TestUserService_UserAlreadyExistsByEmail(t *testing.T) {
	repo := mocks.MockUserRepository{}
	svc := mocks.NewMockUserService(repo)

	testCases := []string{
		"0012345678",
		"99388577322",
		"2335548775889",
		"9999999999",
	}

	for _, tc := range testCases {
		if svc.UserExistsByPhone(tc) {
			t.Logf("User with phone no. %s already exists", tc)
		}
	}
}
func TestUserService_UserAlreadyExistsByPhone(t *testing.T) {
	repo := mocks.MockUserRepository{}
	svc := mocks.NewMockUserService(repo)

	testCases := []string{
		"john.doe@example.com",
		"jane.doe@gmail.com",
		"tony.stark@example.com",
	}

	for _, tc := range testCases {
		if svc.UserExistsByPhone(tc) {
			t.Logf("User with email %s already exists", tc)
		}
	}
}
