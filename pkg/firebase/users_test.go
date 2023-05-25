package firebase

import (
	"testing"

	"firebase.google.com/go/db"
	"github.com/golang/mock/gomock"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/test/mocks"
	"github.com/stretchr/testify/assert"
)

type testUserDatabase struct {
	creator *mocks.MockReferenceCreator
	db      *UserDatabase
}

func setupTestUserDatabase(t *testing.T) testUserDatabase {
	ctrl := gomock.NewController(t)
	creator := mocks.NewMockReferenceCreator(ctrl)
	database := NewUserDatabase(creator)
	return testUserDatabase{
		creator: creator,
		db:      database,
	}
}

func TestGetUser_ValidArgs_ShouldSucceed(t *testing.T) {
	tester := setupTestUserDatabase(t)
	expected := generateRandomUser()
	tester.creator.EXPECT().
		NewRef("users/" + expected.Id).
		Return(db.Ref{}).
		Times(1)

	response, err := tester.db.GetUser(nil, expected.Id)
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
