package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	metrics "github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/api/v1"
	"github.com/readactedworks/go-http-server/pkg/firedb"
	"github.com/readactedworks/go-http-server/pkg/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	tcp                = "tcp"
	port               = 9099
	firebaseDb         = "firebase"
	mysql              = "mysql"
	postgres           = "postgres"
	mongoDb            = "mongodb"
	mongoDbUrlEnvVar   = "MONGO_DB_URL"
	firebaseUrlEnvVar  = "FIREBASE_DB_URL"
	firebaseCredEnvVar = "FIREBASE_CRED_FILEPATH"
)

func main() {
	var err error
	ctx := context.Background()

	server := grpc.NewServer()
	registerUserService(ctx, server, mongoDb)

	//registerServiceMetrics(server, prometheus.DefaultRegisterer)
	reflection.Register(server)
	lis, err := net.Listen(tcp, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerUserService(
	ctx context.Context,
	server *grpc.Server,
	databaseType string,
) {
	switch databaseType {
	case firebaseDb:
		conf := &firebase.Config{DatabaseURL: os.Getenv(firebaseUrlEnvVar)}
		opt := option.WithCredentialsFile(os.Getenv(firebaseCredEnvVar))
		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log.Fatalln("error in initializing firebase app: ", err)
		}

		database, err := app.Database(ctx)
		if err != nil {
			log.Fatalln("error in creating firebase DB client: ", err)
		}
		registerFirebaseUserService(server, database)
	case mongoDb:
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(mongoDbUrlEnvVar)))
		if err != nil {
			log.Fatalln("error in initializing mongo client: ", err)
		}
		registerMongoUserService(server, client)

	default:
		log.Fatalf("database type %s not supported", databaseType)
	}
}

func registerMongoUserService(server *grpc.Server, client *mongo.Client) {
	collection := client.Database("playground").Collection("users")
	user := mongodb.NewUserDatabase(collection)
	svc, err := v1.NewUserService(user, logrus.New())
	if err != nil {
		panic(fmt.Sprintf("user service failed initialization - " + err.Error()))
	}
	model.RegisterUserServiceServer(server, svc)
	logrus.Info("user service registered")
}

func registerFirebaseUserService(server *grpc.Server, fireDb *db.Client) {
	users := firedb.NewUserDatabase(fireDb)
	svc, err := v1.NewUserService(users, logrus.New())
	if err != nil {
		panic(fmt.Sprintf("user service failed initialization - " + err.Error()))
	}

	model.RegisterUserServiceServer(server, svc)
	logrus.Info("user service registered")
}

func registerServiceMetrics(server *grpc.Server, reg prometheus.Registerer) {
	srvMetrics := metrics.NewRegisteredServerMetrics(
		reg,
		metrics.WithServerHandlingTimeHistogram(),
	)
	registry := prometheus.NewRegistry()
	registry.MustRegister(srvMetrics)
	srvMetrics.InitializeMetrics(server)
	logrus.Info("service metrics registered")
}
