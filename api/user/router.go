package user

import (
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", handler.LoginUser)
		r.Post("/signup", handler.CreateUser)
		r.Post("/profile", handler.FetchUser)
	})

	return r
}
