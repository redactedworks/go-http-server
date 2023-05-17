package v1

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type testUserService struct {
	service *UserService
	ctx     context.Context
	manager *mocks.MockUserManager
}

func newTestUserService(t *testing.T) *testUserService {
	ctrl := gomock.NewController(t)
	manager := mocks.NewMockUserManager(ctrl)
	service, _ := NewUserService(manager, logrus.New())

	return &testUserService{
		manager: manager,
		service: service,
		ctx:     context.Background(),
	}
}

func TestGetUser_ValidRequest_ShouldSucceed(t *testing.T) {
	tester := newTestUserService(t)
	expected := generateRandomUser()
	request := &model.GetUserRequest{
		UserId: expected.Id,
	}

	tester.manager.EXPECT().
		GetUser(tester.ctx, expected.Id).
		Return(expected, nil).
		Times(1)

	response, err := tester.service.GetUser(tester.ctx, request)
	assert.NoError(t, err)
	assertUserEqual(t, expected, response.User)
}

func generateRandomUser() *model.User {
	return &model.User{
		Id:       "test-user-id",
		Name:     "test-user-name",
		Email:    "test-user-email",
		Password: "test-user-password",
	}
}

func assertUserEqual(t *testing.T, expected *model.User, actual *model.User) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Password, actual.Password)
}
