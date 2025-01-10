package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	log "github.com/sirupsen/logrus"
)

func CreateAccount(w http.ResponseWriter, r *http.Request){
	var param = api.Users{}
	err := json.NewDecoder(r.Body).Decode(&param)

	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Error("Database connection error: ", err)
        api.InternalErrorHandler(w)
        return
    }

	_, err = database.DB.Exec("INSERT INTO users (name, username, password) VALUES ($1, $2, $3)",
		param.Name, param.Username, param.Password)
	if err != nil {
		log.Error("Failed to insert user: ", err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.Users{
		Name: param.Name,
		Username: param.Username,
		Password: param.Password,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}