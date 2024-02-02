package service

import (
	"context"
	"strings"

	"github.com/Schierke/tasks-handler/internals/models"
)

type TaskRepository interface {
	GetListOfTasks(ctx context.Context, status string) ([]models.TaskResult, error)
}

type taskserviceImpl struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskserviceImpl {
	return &taskserviceImpl{
		repo: repo,
	}
}

func (t *taskserviceImpl) GetTasks(ctx context.Context, status string) ([]models.TaskResult, error) {
	status = strings.ToLower(status)
	return t.repo.GetListOfTasks(ctx, status)
}
