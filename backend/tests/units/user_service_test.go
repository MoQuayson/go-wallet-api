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
