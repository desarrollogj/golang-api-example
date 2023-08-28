package database

import (
	"context"
	"time"

	"github.com/desarrollogj/golang-api-example/libs/logger"
	appConfig "github.com/gookit/config/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *MongoDB

type MongoDB struct {
	Client *mongo.Client
}

func MongoConnect() *MongoDB {
	// connect
	opts := options.Client().ApplyURI(appConfig.String("database.connectionString"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.AppLog.Fatal().Err(err).Msg("unable to create a Connection")
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.AppLog.Fatal().Err(err).Msg("unable to ping database")
	}

	logger.AppLog.Info().Msg("connected to MongoDB")

	return &MongoDB{Client: client}
}

func MongoDisconnect() {
	logger.AppLog.Info().Msg("disconnecting to MongoDB")

	if Mongo == nil {
		return
	}

	if err := Mongo.Client.Disconnect(context.TODO()); err != nil {
		logger.AppLog.Fatal().Err(err).Msg("unable to close connection")
	}

	logger.AppLog.Info().Msg("disconnected to MongoDB")
}
