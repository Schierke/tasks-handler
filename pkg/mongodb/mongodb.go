package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/Schierke/tasks-handler/config"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connString    = "mongodb://%s:%s"
	defautTimeout = 10 * time.Second
	datasetDir    = "./db/datasets/"
)

func SetupDB(cfg config.AppConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defautTimeout)
	defer cancel()

	dbUrl := fmt.Sprintf(connString,
		cfg.Mongo.Host,
		cfg.Mongo.Port)

	credential := options.Credential{
		Username: cfg.Mongo.User,
		Password: cfg.Mongo.Password,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl).SetAuth(credential))

	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to mongo DB")
		return nil, err
	}

	return client, nil
}
