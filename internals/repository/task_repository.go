package repository

import (
	"context"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
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

func (t *taskRepo) GetListOfTasks(ctx context.Context, opts ...func(*Pipeline)) ([]models.TaskResult, error) {
	var result []models.TaskResult

	pipeline := make(Pipeline, 0)

	for _, opt := range opts {
		opt(&pipeline)
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

func (t *taskRepo) FindTaskById(ctx context.Context, id string) (*models.Task, error) {
	var (
		task models.Task
		err  error
	)

	err = t.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)

	if err != nil {
		return &task, err
	}

	return &task, nil
}

func (t *taskRepo) UpdateTaskAssignee(ctx context.Context, taskId, assigneeId string) error {
	filter := bson.D{{Key: "_id", Value: taskId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "assigneeId", Value: assigneeId}}}}

	_, err := t.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) GetListOfTasksWithOpsMemberName(ctx context.Context, opts ...func(*Pipeline)) ([]models.TaskResult2, error) {
	var result []models.TaskResult2

	pipeline := make(Pipeline, 0)

	for _, opt := range opts {
		opt(&pipeline)
	}
	pipeline = append(pipeline, createListTasks2())

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
