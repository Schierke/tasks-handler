package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type TaskService interface {
	GetTasks(ctx context.Context, status string) ([]models.TaskResult, error)
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
