package main

import (
	"context"
	"fmt"

	"firebase.google.com/go/db"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/api/v1"
	"github.com/readactedworks/go-http-server/pkg/firebase"
	"google.golang.org/grpc"
)

func main() {
	database, err := db.NewClient(context.Background(), nil)
	if err != nil {
		panic("")
	}

	server := grpc.NewServer()
	registerUserService(server, database)
	registerCompanyService(server, database)
}

func registerUserService(server *grpc.Server, database *db.Client) {
	users := firebase.NewUserDatabase(database)
	svc, err := v1.NewUserService(users, nil)
	if err != nil {
		panic(fmt.Sprintf("user service failed initialization - " + err.Error()))
	}

	model.RegisterUserServiceServer(server, svc)
}

func registerCompanyService(server *grpc.Server, database *db.Client) {
}
