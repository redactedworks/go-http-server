package v1

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/readactedworks/go-http-server/api/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	notFound     = "not found"
	userLogField = "user_id"
)

var ErrUserIdMissing = errors.New("user id was not specified")
var ErrUserEmailMissing = errors.New("user email was not specified")
var ErrUserEmailInvalid = errors.New("user email was invalid")
var ErrUserNameMissing = errors.New("user name was not specified")
var ErrUserDeletionFailed = errors.New("user deletion failed")
var ErrUserPasswordMissing = errors.New("user password was not specified")

// UserDataManager provides basic CRUD database operations for Users.
type UserDataManager interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}

// UserService provides functionality to manage Users.
type UserService struct {
	model.UnimplementedUserServiceServer

	db  UserDataManager
	log *logrus.Logger
}

// NewUserService creates a new instance of a UserService.
func NewUserService(
	db UserDataManager,
	log *logrus.Logger,
) (*UserService, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	if log == nil {
		return nil, errors.New("log is required")
	}
	return &UserService{
		db:  db,
		log: log,
	}, nil
}

// GetUser retrieves a User by their ID from the database.
func (s *UserService) GetUser(
	ctx context.Context,
	request *model.GetUserRequest,
) (*model.GetUserResponse, error) {
	if err := isValidGetUserRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.db.GetUser(ctx, request.UserId)
	if err != nil {
		s.log.
			WithField(userLogField, request.UserId).
			Error(err)

		if strings.Contains(err.Error(), notFound) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &model.GetUserResponse{
		User: user,
	}, nil
}

// CreateUser generates a new User in the database.
func (s *UserService) CreateUser(
	ctx context.Context,
	request *model.CreateUserRequest,
) (*model.CreateUserResponse, error) {
	if err := isValidCreateUserRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := &model.User{
		Name:     strings.TrimSpace(request.Name),
		Email:    strings.TrimSpace(request.Email),
		Password: strings.TrimSpace(request.Password),
	}

	if err := s.db.CreateUser(ctx, user); err != nil {
		s.log.
			WithField(userLogField, user.Id).
			Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &model.CreateUserResponse{Success: true}, nil
}

// UpdateUser modifies an existing User in the database.
func (s *UserService) UpdateUser(
	ctx context.Context,
	request *model.UpdateUserRequest,
) (*model.UpdateUserResponse, error) {
	if err := isValidUpdateUserRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := &model.User{
		Id:       strings.TrimSpace(request.Id),
		Name:     strings.TrimSpace(request.Name),
		Email:    strings.TrimSpace(request.Email),
		Password: strings.TrimSpace(request.Password),
	}

	if err := s.db.UpdateUser(ctx, user); err != nil {
		s.log.
			WithField(userLogField, user.Id).
			Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &model.UpdateUserResponse{
		Updated: true,
	}, nil
}

// DeleteUser removes an existing User from the database.
func (s *UserService) DeleteUser(
	ctx context.Context,
	request *model.DeleteUserRequest,
) (*model.DeleteUserResponse, error) {
	if err := isValidDeleteUserRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.db.DeleteUser(ctx, request.UserId); err != nil {
		s.log.
			WithField(userLogField, request.UserId).
			Error(err)
		return nil, status.Error(codes.Internal, ErrUserDeletionFailed.Error())
	}

	return &model.DeleteUserResponse{Deleted: true}, nil
}

func isValidGetUserRequest(request *model.GetUserRequest) error {
	if strings.TrimSpace(request.UserId) == "" {
		return ErrUserIdMissing
	}

	return nil
}

func isValidCreateUserRequest(request *model.CreateUserRequest) error {
	if strings.TrimSpace(request.Name) == "" {
		return ErrUserNameMissing
	}
	if strings.TrimSpace(request.Email) == "" {
		return ErrUserEmailMissing
	}
	if _, err := mail.ParseAddress(
		strings.TrimSpace(request.Email),
	); err != nil {
		return ErrUserEmailInvalid
	}
	if strings.TrimSpace(request.Password) == "" {
		return ErrUserPasswordMissing
	}

	return nil
}

func isValidUpdateUserRequest(request *model.UpdateUserRequest) error {
	if strings.TrimSpace(request.Id) == "" {
		return ErrUserIdMissing
	}
	if strings.TrimSpace(request.Name) == "" {
		return ErrUserNameMissing
	}
	if strings.TrimSpace(request.Email) == "" {
		return ErrUserEmailMissing
	}
	if _, err := mail.ParseAddress(
		strings.TrimSpace(request.Email),
	); err != nil {
		return ErrUserEmailInvalid
	}
	if strings.TrimSpace(request.Password) == "" {
		return ErrUserPasswordMissing
	}
	return nil
}

func isValidDeleteUserRequest(request *model.DeleteUserRequest) error {
	if strings.TrimSpace(request.UserId) == "" {
		return ErrUserIdMissing
	}

	return nil
}
