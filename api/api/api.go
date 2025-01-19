package api

import (
	"encoding/json"
	"net/http"
)

type Users struct {
    Name string `json:"name"`
	Email string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Token_validity struct {
	Message string
}

type Login_user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Error struct {
	Code    int
	Message string
}

type User_Create_Response struct {
	Message string
	Token string
}

func writeError(w http.ResponseWriter, message string, code int) {
    resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
	ClientErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Invalid Request", http.StatusBadRequest)
	}
)