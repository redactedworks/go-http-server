package firebase

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

// ReferenceManager todo:describe
type ReferenceManager interface {
	Get(ctx context.Context, v interface{}) error
	Set(ctx context.Context, v interface{}) error
	Delete(ctx context.Context) error
	Update(ctx context.Context, v map[string]interface{}) error
}

// Referencer todo:describe
type Referencer interface {
	NewRef(path string) *db.Ref
}

// UserDatabase todo:describe
type UserDatabase struct {
	db Referencer
}

// NewUserDatabase todo:describe
func NewUserDatabase(db Referencer) *UserDatabase {
	return &UserDatabase{
		db: db,
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

	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, id))
	var user model.User
	if err := ref.Get(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser todo:describe
func (u *UserDatabase) CreateUser(ctx context.Context, user *model.User) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// UpdateUser todo:describe
func (u *UserDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, user.Id))
	if err := ref.Set(ctx, &user); err != nil {
		return err
	}

	return nil
}

// DeleteUser todo:describe
func (u *UserDatabase) DeleteUser(ctx context.Context, id string) error {
	ref := u.db.NewRef(fmt.Sprintf(userReferenceFmt, id))
	if err := ref.Delete(ctx); err != nil {
		return err
	}

	return nil
}
