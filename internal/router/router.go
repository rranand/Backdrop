package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	custommiddleware "github.com/rranand/backdrop/internal/middleware"
	"github.com/rranand/backdrop/user"
)

type Route struct {
	Path   string
	Router *chi.Mux
}

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(custommiddleware.JsonMiddleware)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routes := []Route{
		{Path: "/auth", Router: user.AuthRouter()},
	}

	for _, route := range routes {
		r.Mount(route.Path, route.Router)
	}

	return r
}
