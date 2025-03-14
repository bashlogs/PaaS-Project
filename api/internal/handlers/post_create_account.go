package handlers

import (
	"encoding/json"
	"errors"
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

	// check for the invalid request
	var email string
	err = database.DB.QueryRow("select username from users where email=$1", param.Email).Scan(&email)
	if err == nil {
		log.Error("Email already exists")
		api.RequestErrorHandler(w, errors.New("email already exists"))
		return
	}

	var username string
	err = database.DB.QueryRow("select username from users where username=$1", param.Username).Scan(&username)
	if err == nil {
		log.Error("Username already exists")
		api.RequestErrorHandler(w, errors.New("username already exists"))
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)", param.Name, param.Username, param.Email, param.Password)
	if err != nil {
		log.Error("Failed to insert user: ", err)
		api.InternalErrorHandler(w)
		return
	}

	token, err := middleware.JWT_Token(param.Email)
	if err != nil {
		log.Error("Failed to generate token: ", err)
		api.InternalErrorHandler(w)
		return
	}
	
	// http.SetCookie(w, &http.Cookie{	
	// 	Name: "token",
	// 	Value: token,
	// 	HttpOnly: true,
	// 	MaxAge: 7200,
	// 	Path: "/",
	// })
	// log.Println("Set-Cookie: token=", token)

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
	
	// fmt.Println("param.Username:", param.Username, "param.Password", param.Password)
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
		token, err := middleware.JWT_Token(param.Username)
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

type DashboardResponse struct {
	Message string `json:"message"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

var SecretKey = []byte("khadde")

// func Dashboard(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve the auth token from the cookie
// 	cookie, err := r.Cookie("authToken")
// 	if err != nil {
// 		log.Error("Cookie error:", err)
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	tokenString := cookie.Value
// 	if tokenString == "" {
// 		log.Error("Authorization token missing")
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	// Parse and validate the JWT token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Ensure the signing method is HMAC
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return SecretKey, nil // Replace secretKey with your actual secret
// 	})

// 	if err != nil || !token.Valid {
// 		log.Error("Invalid token:", err)
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Extract claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		log.Error("Failed to parse token claims")
// 		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
// 		return
// 	}

// 	// Extract email from token
// 	email, ok := claims["email"].(string)
// 	if !ok {
// 		log.Error("Invalid email in token")
// 		http.Error(w, "Invalid token: missing email", http.StatusUnauthorized)
// 		return
// 	}

// 	fmt.Println("Email:", email)
// 	// Connect to the database
// 	database, err := tools.ConnectToDatabase()
// 	if err != nil {
// 		log.Error("Database connection error:", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	var username, name string
// 	err = database.DB.QueryRow("SELECT username, name FROM users WHERE email=$1", email).Scan(&username, &name)
// 	if err != nil {
// 		log.Error("Error retrieving user details:", err)
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	// Prepare the response
// 	response := DashboardResponse{
// 		Message: "Access granted to the dashboard",
// 		Email:    email,
// 		Username: username,
// 		Name:     name,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(response)
// 	if err != nil {
// 		log.Error("Error encoding dashboard response:", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 	}
// }


func Dashboard(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "Failed to retrieve user information", http.StatusUnauthorized)
		return
	}

	database, err := tools.ConnectToDatabase()
	if err != nil {
		log.Error("Database connection error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var username, name string 
	err = database.DB.QueryRow("SELECT username, name FROM users WHERE email=$1", email).Scan(&username, &name)
	if err != nil {
		log.Error("Error retrieving user data:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := DashboardResponse{
		Message: "Access granted to the dashboard",
		Email:    email,
		Username: username,
		Name:     name,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Error encoding dashboard response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	// fmt.Fprintf(w, "Welcome %s! Your data: %v", email, userData)
}
