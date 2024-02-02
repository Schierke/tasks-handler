package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Schierke/tasks-handler/internals/models"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type ShiftService interface {
	GetShifts(context.Context, string) ([]models.ShiftResponse, error)
}

type shiftHandler struct {
	service ShiftService
}

func NewShiftHandler(service ShiftService, r *chi.Mux) {
	handler := &shiftHandler{
		service: service,
	}
	r.Route("/api/v1/shifts", func(r chi.Router) {
		r.Get("/", handler.GetShifts)
	})
}

func (s *shiftHandler) GetShifts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskid := r.URL.Query().Get("taskid")

	if taskid == "" {
		log.Error().Msg(ErrTaskIdNotFound)
		http.Error(w, ErrTaskIdNotFound, http.StatusBadRequest)
		return
	}

	shifts, err := s.service.GetShifts(ctx, taskid)

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

	if err := json.NewEncoder(w).Encode(shifts); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
