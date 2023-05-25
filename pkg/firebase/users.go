package firebase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/pkg/firebase/references"
)

const (
	userReferenceFmt = "users/%s"
)

// UserDatabase provides access to User-specific actions in Firebase real-time
// data store.
type UserDatabase struct {
	referenceCreator references.Creator
}

// NewUserDatabase creates a new instance of UserDatabase.
func NewUserDatabase(db references.Creator) *UserDatabase {
	return &UserDatabase{
		referenceCreator: db,
	}
}

// GetUser todo:describe
func (u *UserDatabase) GetUser(
	ctx context.Context,
	id string,
) (*model.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("id was missing")
	}

	entry := fmt.Sprintf(userReferenceFmt, id)
	ref := u.referenceCreator.NewRef(entry)
	var user model.User
	if err := ref.Get(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser todo:describe
func (u *UserDatabase) CreateUser(ctx context.Context, user *model.User) error {
	entry := fmt.Sprintf(userReferenceFmt, user.Id)
	ref := u.referenceCreator.NewRef(entry)
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// UpdateUser todo:describe
func (u *UserDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	entry := fmt.Sprintf(userReferenceFmt, user.Id)
	ref := u.referenceCreator.NewRef(entry)
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// DeleteUser todo:describe
func (u *UserDatabase) DeleteUser(ctx context.Context, id string) error {
	entry := fmt.Sprintf(userReferenceFmt, id)
	ref := u.referenceCreator.NewRef(entry)
	if err := ref.Delete(ctx); err != nil {
		return err
	}

	return nil
}
