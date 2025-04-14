package task

import (
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/new", handler.CreateTask)
	})

	return r
}
