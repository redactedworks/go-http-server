package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/db"
	"github.com/readactedworks/go-http-server/api/model"
)

const (
	userReferenceFmt = "users/%s"
)

type Referencer interface {
	NewRef(path string) *db.Ref
}

type UserDatabase struct {
	db Referencer
}

func NewUserDatabase(db Referencer) *UserDatabase {
	return &UserDatabase{
		db: db,
	}
}

func (u *UserDatabase) GetUser(
	ctx context.Context,
	id string,
) (*model.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("id was missing")
	}

	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, id))
	var user model.User
	if err := ref.Get(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDatabase) CreateUser(ctx context.Context, user *model.User) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) DeleteUser(ctx context.Context, id string) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, id))
	if err := ref.Delete(ctx); err != nil {
		return err
	}
	return nil
}
