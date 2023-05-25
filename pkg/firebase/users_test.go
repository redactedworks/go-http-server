package firebase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/internal/test/mocks"
	"github.com/readactedworks/go-http-server/internal/test/utils"
	"github.com/stretchr/testify/assert"
)

type testUserDatabase struct {
	ctrl     *gomock.Controller
	creator  *mocks.MockReferenceCreator
	operator *mocks.MockReferenceOperator
	db       *UserDatabase
	ctx      context.Context
}

func setupTestUserDatabase(t *testing.T) testUserDatabase {
	ctrl := gomock.NewController(t)
	creator := mocks.NewMockReferenceCreator(ctrl)
	operator := mocks.NewMockReferenceOperator(ctrl)
	database := NewUserDatabase(creator)

	return testUserDatabase{
		creator:  creator,
		db:       database,
		operator: operator,
		ctrl:     ctrl,
		ctx:      context.Background(),
	}
}

func TestGetUser_ValidArgs_ShouldSucceed(t *testing.T) {
	tester := setupTestUserDatabase(t)
	expected := utils.GenerateRandomUser()

	var actual model.User
	tester.operator.EXPECT().
		Get(tester.ctx, &actual).
		DoAndReturn(func(ctx context.Context, value any) error {
			*value.(*model.User) = *expected
			return nil
		}).
		Times(1)

	tester.creator.EXPECT().
		NewRef("users/" + expected.Id).
		Return(tester.operator).
		Times(1)

	response, err := tester.db.GetUser(tester.ctx, expected.Id)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func generateRandomUser() *model.User {
	return &model.User{
		Id:       "1",
		Name:     "test",
		Email:    "",
		Password: "",
	}
}
