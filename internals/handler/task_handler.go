package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type TaskService interface {
	GetTasks(ctx context.Context, status string) ([]models.TaskResult, error)
	SetAssignedOpsMember(ctx context.Context, taskId, assigneeId string) error
	GetTaskLocation(ctx context.Context, taskId string) (*models.Address, error)
	GetTasksWithOps(ctx context.Context) ([]models.TaskResult2, error)
}

type taskHandler struct {
	service TaskService
}

func NewTaskHandler(service TaskService, r *chi.Mux) {
	handler := &taskHandler{
		service: service,
	}
	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Get("/", handler.GetTasks)
		r.Get("/ops", handler.GetTasksWithOps)
		r.Route("/{taskID}", func(r chi.Router) {
			r.Patch("/update", handler.SetAssignedOpsMember)
			r.Get("/location", handler.GetTaskLocation)
		})
	})
}

func (t *taskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	shiftStatus := r.URL.Query().Get("status")
	taskResults, err := t.service.GetTasks(ctx, shiftStatus)

	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, Timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, ErrTasksNotFound.Error(), http.StatusInternalServerError)

		}
		return
	}

	// Send an OK status
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(taskResults); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *taskHandler) SetAssignedOpsMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskId := chi.URLParam(r, "taskID")
	if taskId == "" {
		log.Fatal().Err(errors.New(ErrTaskIdNotFound))
		http.Error(w, ErrTaskIdNotFound, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't ready body when create new project")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type input struct {
		AssigneeId string `json:"assignee_id"`
	}
	var assignee input
	json.Unmarshal(body, &assignee)

	err = t.service.SetAssignedOpsMember(ctx, taskId, assignee.AssigneeId)

	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, Timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, ErrTasksNotFound.Error(), http.StatusInternalServerError)

		}
		return
	}

	// Send an OK status
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (t *taskHandler) GetTaskLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskId := chi.URLParam(r, "taskID")
	if taskId == "" {
		log.Fatal().Err(errors.New(ErrTaskIdNotFound))
		http.Error(w, ErrTaskIdNotFound, http.StatusBadRequest)
		return

	}

	address, err := t.service.GetTaskLocation(ctx, taskId)

	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, Timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, ErrLocationNotFound, http.StatusInternalServerError)

		}
		return
	}

	// Send an OK status
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(address); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (t *taskHandler) GetTasksWithOps(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskResults, err := t.service.GetTasksWithOps(ctx)

	if err != nil {
		log.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, Timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, ErrTasksNotFound.Error(), http.StatusInternalServerError)

		}
		return
	}

	// Send an OK status
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(taskResults); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
