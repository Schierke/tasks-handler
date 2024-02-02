package service

import (
	"context"
	"strings"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/Schierke/tasks-handler/internals/repository"
)

type TaskRepository interface {
	GetListOfTasks(ctx context.Context, opts ...func(*repository.Pipeline)) ([]models.TaskResult, error)
	FindTaskById(ctx context.Context, id string) (*models.Task, error)
	UpdateTaskAssignee(ctx context.Context, task, assigneeId string) error
	GetListOfTasksWithOpsMemberName(ctx context.Context, opts ...func(*repository.Pipeline)) ([]models.TaskResult2, error)
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
	return t.repo.GetListOfTasks(ctx, repository.WithOrganisationLookup(), repository.WithSlotLookUp(), repository.WithFilterShiftStatus(status))
}

func (t *taskserviceImpl) SetAssignedOpsMember(ctx context.Context, taskId, assigneeId string) error {
	task, err := t.repo.FindTaskById(ctx, taskId)

	if err != nil {
		return err
	}

	t.repo.UpdateTaskAssignee(ctx, task.ID, assigneeId)

	return nil
}

func (t *taskserviceImpl) GetTaskLocation(ctx context.Context, taskId string) (*models.Address, error) {
	task, err := t.repo.FindTaskById(ctx, taskId)
	if err != nil {
		return nil, err
	}

	address := &task.Address

	return address, nil
}

func (t *taskserviceImpl) GetTasksWithOps(ctx context.Context) ([]models.TaskResult2, error) {
	return t.repo.GetListOfTasksWithOpsMemberName(ctx, repository.WithUsersLookup())
}
