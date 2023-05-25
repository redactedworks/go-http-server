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
	referenceCreator  references.Creator
	referenceOperator references.OperatorCreator
}

// NewUserDatabase creates a new instance of UserDatabase.
func NewUserDatabase(db references.Creator) *UserDatabase {
	return &UserDatabase{
		referenceCreator:  db,
		referenceOperator: &references.Factory{},
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

	ref := u.referenceCreator.NewRef(fmt.Sprintf(userReferenceFmt, id))
	operator := u.referenceOperator.NewOperator(ref)
	var user model.User
	if err := operator.Get(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser todo:describe
func (u *UserDatabase) CreateUser(ctx context.Context, user *model.User) error {
	ref := u.referenceCreator.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	operator := u.referenceOperator.NewOperator(ref)
	if err := operator.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// UpdateUser todo:describe
func (u *UserDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	ref := u.referenceCreator.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	operator := u.referenceOperator.NewOperator(ref)
	if err := operator.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// DeleteUser todo:describe
func (u *UserDatabase) DeleteUser(ctx context.Context, id string) error {
	ref := u.referenceCreator.NewRef(fmt.Sprintf(userReferenceFmt, id))
	operator := u.referenceOperator.NewOperator(ref)
	if err := operator.Delete(ctx); err != nil {
		return err
	}

	return nil
}
