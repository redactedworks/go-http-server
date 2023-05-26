package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	metrics "github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/readactedworks/go-http-server/api/model"
	"github.com/readactedworks/go-http-server/api/v1"
	"github.com/readactedworks/go-http-server/pkg/firedb"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const (
	dbUrl     = ""
	credsFile = "config/go-backend-firebase.json"
)

func main() {
	ctx := context.Background()
	conf := &firebase.Config{DatabaseURL: dbUrl}
	opt := option.WithCredentialsFile(credsFile)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	database, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	server := grpc.NewServer()
	registerUserService(server, database)
	//registerServiceMetrics(server, prometheus.DefaultRegisterer)
}

func registerUserService(server *grpc.Server, database *db.Client) {
	users := firedb.NewUserDatabase(database)
	svc, err := v1.NewUserService(users, nil)
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
