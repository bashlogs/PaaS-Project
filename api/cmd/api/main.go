package main

import (
	"fmt"
	"net/http"

	"github.com/bashlogs/PaaS_Project/api/internal/handlers"
	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	
	fmt.Println("Starting Kubernetes API Service...")

	fmt.Println(`
			
██   ██ ██    ██ ██████  ███████ ██████  ███    ██ ███████ ████████ ███████ ███████      █████  ██████  ██ 
██  ██  ██    ██ ██   ██ ██      ██   ██ ████   ██ ██         ██    ██      ██          ██   ██ ██   ██ ██ 
█████   ██    ██ ██████  █████   ██████  ██ ██  ██ █████      ██    █████   ███████     ███████ ██████  ██ 
██  ██  ██    ██ ██   ██ ██      ██   ██ ██  ██ ██ ██         ██    ██           ██     ██   ██ ██      ██ 
██   ██  ██████  ██████  ███████ ██   ██ ██   ████ ███████    ██    ███████ ███████     ██   ██ ██      ██ 
																										
	`)

	err := http.ListenAndServe(":8000", enableCORS(r))
	if err != nil {
		log.Error(err)
	}


	// Testcases

	// CreateTable()
	// Print()
}