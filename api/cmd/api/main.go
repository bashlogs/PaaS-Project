package main

import (
	"fmt"
	"net/http"

	"github.com/bashlogs/PaaS_Project/api/internal/handlers"
	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"
)

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

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}


	// Testcases

	// CreateTable()
	// Print()
}