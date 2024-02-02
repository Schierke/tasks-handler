package repository

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	shiftCollection = "shifts"
)

type shiftRepo struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func NewtShiftRepo(DB *mongo.Database) *shiftRepo {
	return &shiftRepo{
		DB:         DB,
		Collection: DB.Collection(shiftCollection),
	}
}

func (t *shiftRepo) GetShiftsByTaskID(ctx context.Context, id string) ([]models.Shift, error) {
	var (
		shifts []models.Shift
		err    error
	)

	cursor, err := t.Collection.Find(ctx, bson.M{"taskId": id})

	if err != nil {
		log.Fatal().Err(err).Msg("can't find shifts with given task id")
		return nil, err
	}

	err = cursor.All(ctx, &shifts)
	if err != nil {
		log.Fatal().Err(err).Msg("can't find shifts with given task id")
		return nil, err
	}

	return shifts, nil
}
