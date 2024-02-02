package repository

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	slotCollection = "slots"
)

type slotRepo struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func NewtSlotRepo(DB *mongo.Database) *slotRepo {
	return &slotRepo{
		DB:         DB,
		Collection: DB.Collection(slotCollection),
	}
}

func (s *slotRepo) GetSlotsByShiftID(ctx context.Context, id string) ([]models.Slot, error) {
	var (
		slots []models.Slot
		err   error
	)

	cursor, err := s.Collection.Find(ctx, bson.M{"shiftId": id})

	if err != nil {
		log.Fatal().Err(err).Msg("can't find slots with given shift id")
		return nil, err
	}

	err = cursor.All(ctx, &slots)
	if err != nil {
		log.Fatal().Err(err).Msg("can't find slots with given shift id")
		return nil, err
	}

	return slots, nil
}
