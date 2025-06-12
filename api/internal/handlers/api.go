package handlers

import (
	"github.com/bashlogs/PaaS_Project/api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Handler(r *chi.Mux) {
    r.Use(chimiddle.StripSlashes)

	r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8000"}, // Allow your frontend origin
        AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}, // Include OPTIONS
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
    }))
   
    r.Post("/signup", CreateAccount)
    r.Post("/login", Login)
    r.Route("/dashboard", func(router chi.Router){
        router.Use(middleware.Authorization)
        router.Get("/", Dashboard)
    })

    r.Route("/api", func(router chi.Router){
        router.Use(middleware.Authorization)
        router.Get("/workspaces", GetWorkspaces)
        router.Post("/workspaces", CreateWorkspace)
        router.Delete("/workspaces", DeleteWorkspace)
        router.Post("/workspaces_status", UpdateWorkspace)
    })

    r.Route("/deployment", func(router chi.Router){
        router.Use(middleware.Authorization)
        router.Post("/create", CreateDeployment2)
        // router.Get("/status", GetDeploymentStatus)
        // router.Post("/create", CreateDeployment)
        // router.Delete("/delete", DeleteDeployment)
    })
}