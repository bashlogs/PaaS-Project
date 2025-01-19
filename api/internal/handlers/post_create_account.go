package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/middleware"
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

	token, err := middleware.JWT_Token(param.Username, param.Password)
	if err != nil {
		log.Error("Failed to generate token: ", err)
		api.InternalErrorHandler(w)
		return
	}
	
	http.SetCookie(w, &http.Cookie{	
		Name: "token",
		Value: token,
		HttpOnly: true,
		MaxAge: 7200,
		Path: "/users",
	})
	log.Println("Set-Cookie: token=", token)


	var response = api.User_Create_Response{
		Message: "Successfully Logined",
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	var param = api.Login_user{}
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
	
	fmt.Println("param.Username:", param.Username, "param.Password", param.Password)
	var password string
	err = database.DB.QueryRow("select password from users where username=$1 OR email=$1", param.Username).Scan(&password)

	if err == sql.ErrNoRows {
		if err == sql.ErrNoRows {
			log.Error("user/email not found")
			api.RequestErrorHandler(w, errors.New("user/email not found"))
			return
		}
		log.Error("database error: ", err)
		api.InternalErrorHandler(w)
		return
	}

	if password == param.Password {
		token, err := middleware.JWT_Token(param.Username, param.Password)
		if err != nil {
			log.Error("Failed to generate token: ", err)
			api.InternalErrorHandler(w)
			return
		}

		var response = api.User_Create_Response{
			Message: "Successfully Logined",
			Token: token,
		}
	
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
	} else{
		api.RequestErrorHandler(w, errors.New("wrong Password"))
		return
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Dashboard"))
}