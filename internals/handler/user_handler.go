package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserService interface {
}

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService, r *chi.Mux) {
	handler := &userHandler{
		service: service,
	}
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", nil)
		r.Post("/register", handler.SetAssignedOpsMember)
	})
}

func (u *userHandler) SetAssignedOpsMember(w http.ResponseWriter, r *http.Request) {

}
