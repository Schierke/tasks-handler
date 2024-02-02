package repository

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	taskCollection = "tasks"
)

type taskRepo struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func NewTaskRepo(DB *mongo.Database) *taskRepo {
	return &taskRepo{
		DB:         DB,
		Collection: DB.Collection(taskCollection),
	}
}

func (t *taskRepo) GetListOfTasks(ctx context.Context, status string) ([]models.TaskResult, error) {
	var result []models.TaskResult

	pipeline := make(mongo.Pipeline, 0)

	pipeline = append(pipeline, organisationLookup())
	pipeline = append(pipeline, slotLookup())
	pipeline = append(pipeline, bson.D{{Key: "$unwind", Value: "$organisation"}})

	if status != "" {
		pipeline = append(pipeline, filterShiftStatus(status))
	}
	pipeline = append(pipeline, createListTasks())

	cursor, err := t.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error().Err(err).Msg("can't launch get tasks pipeline")
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &result); err != nil {
		log.Error().Err(err).Msg("can't decode the result")
		return nil, err
	}

	return result, nil
}

func (t *taskRepo) FindOne(ctx context.Context, id string) (*models.Task, error) {
	var (
		task models.Task
		err  error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = t.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&task)
	
	if err != nil {
		return &task, err
	}

	return &task, nil
}
