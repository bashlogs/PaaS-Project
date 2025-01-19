package handlers

import (
	"github.com/bashlogs/PaaS_Project/api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {

	r.Use(chimiddle.StripSlashes)

	r.Post("/signup", CreateAccount)
	r.Post("/login", Login)
	
	r.Route("/dashboard", func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/", Dashboard)
	})
}