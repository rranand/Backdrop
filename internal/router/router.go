package router

import (
	"github.com/go-chi/chi"
	"github.com/rranand/backdrop/user"
)

type Route struct {
	Path   string
	Router *chi.Mux
}

func Router() *chi.Mux {
	r := chi.NewRouter()

	routes := []Route{
		{Path: "/auth", Router: user.AuthRouter()},
	}

	for _, route := range routes {
		r.Mount(route.Path, route.Router)
	}

	return r
}
