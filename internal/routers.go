package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type AppRouter struct{}

func NewAppRouter() AppRouter {
	return AppRouter{}
}

func (a AppRouter) Route() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(api chi.Router) {
		api.Route("/invoices", func(inv chi.Router) {
			//inv.Get("/",)
		})
	})

	return r
}
