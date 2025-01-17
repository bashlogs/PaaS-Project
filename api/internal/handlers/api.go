package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {

	r.Use(chimiddle.StripSlashes)

	r.Route("/create_account", func(router chi.Router){
		router.Post("/account", CreateAccount)
	})
}