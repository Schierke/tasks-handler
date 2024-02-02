package migrate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Schierke/tasks-handler/config"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connString = "mongodb://%s:%s"
	datasetDir = "./db/datasets/"
)

func Migrate(config config.AppConfig) {
	dbUrl := fmt.Sprintf(connString,
		config.Mongo.Host,
		config.Mongo.Port)

	credential := options.Credential{
		Username: config.Mongo.User,
		Password: config.Mongo.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl).SetAuth(credential))

	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to mongo DB")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal().Err(err).Msg("error while disconnecting:")
			panic(err)
		}
	}()

	err = migratingDataset(ctx, client, config.Mongo.Name)

	if err != nil {
		log.Fatal().Err(err).Msg("can't migrating dataset")
	}
}

func migratingDataset(ctx context.Context, client *mongo.Client, database string) error {
	files, err := os.ReadDir(datasetDir)
	if err != nil {
		log.Fatal().Err(err).Msg("can't read folder dataset")
		return err
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".json") {
			filePath := filepath.Join(datasetDir, name)

			// Read the JSON file
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal().Err(err).Msg(fmt.Sprintf("can't read dataset: %s", name))
				return err
			}

			collectionName := strings.TrimSuffix(name, ".json")
			collection := client.Database(database).Collection(collectionName)

			// Unmarshal JSON data into a slice of User structs
			var data []interface{}
			if err := json.Unmarshal(fileData, &data); err != nil {
				return err
			}

			// Insert all users into the MongoDB collection
			insertResult, err := collection.InsertMany(ctx, data)
			if err != nil {
				log.Fatal().Err(err).Msg("can't insert to collection")
			}

			fmt.Printf("Inserted %v documents into the %s collection!\n", len(insertResult.InsertedIDs), collectionName)

		}
	}

	return nil
}
