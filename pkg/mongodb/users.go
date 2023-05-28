package mongodb

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/readactedworks/go-http-server/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userId = "_id"
)

type MongoDataStorer interface {
	InsertOne(
		ctx context.Context,
		document interface{},
		opts ...*options.InsertOneOptions,
	) (*mongo.InsertOneResult, error)
	FindOne(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOneOptions,
	) *mongo.SingleResult
	UpdateOne(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
}

type UserDatabase struct {
	collection *mongo.Collection
}

func NewUserDatabase(collection *mongo.Collection) *UserDatabase {
	return &UserDatabase{
		collection: collection,
	}
}

// GetUser retrieves a user from the Mongo database.
func (u *UserDatabase) GetUser(
	ctx context.Context,
	id string,
) (*model.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("id was missing")
	}

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.D{{userId, objectId}}
	var user model.User
	err = u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDatabase) CreateUser(ctx context.Context, user *model.User) error {
	_, err := u.collection.
		InsertOne(ctx, generateDocumentFromUser(user))
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	filter := bson.D{{userId, user.Id}}
	update := bson.D{{"$set", generateDocumentFromUser(user)}}
	_, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func generateDocumentFromUser(user *model.User) bson.D {
	return bson.D{
		{userId, user.Id},
		{"name", user.Name},
		{"email", user.Email},
		{"password", user.Password},
	}
}
