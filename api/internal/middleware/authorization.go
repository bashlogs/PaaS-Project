package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("invalid username or tokens")
var ErrInvalidToken = errors.New("invalid token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var err error
		cookie, err := r.Cookie("token")
		// fmt.Println(cookie)
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Token cookie missing", http.StatusUnauthorized)
				return
			}
			log.Println("Error retrieving cookie:", err)
			http.Error(w, "Invalid cookie", http.StatusBadRequest)
			return
		}
		tokenString := cookie.Value
		if tokenString == "" {
			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
			return
		}
		database, err := tools.ConnectToDatabase()
		if err != nil {
			log.Println("Database connection error:", err)
			api.InternalErrorHandler(w)
			return
		}
	
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			log.Println("Invalid token:", err)
			api.RequestErrorHandler(w, ErrInvalidToken)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Verify token expiration
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					log.Error("Failed to generate token: ", err)
					api.RequestErrorHandler(w, errors.New("token has expired"))
					return
				}
			} else {
				log.Error("Invali exp date: ")
				api.RequestErrorHandler(w, errors.New("invalid expiration in token"))
				return 
			}
			username, ok := claims["username"].(string)
			if !ok {
				log.Error("Failed to insert user1: ", err)
				api.RequestErrorHandler(w, errors.New("invalid username"))
				return 
			}
			var password string
			err = database.DB.QueryRow("select password from users where username=$1", username).Scan(&password)
			if err != nil {
				log.Error("Failed to insert user2: ", err)
				api.RequestErrorHandler(w, errors.New("invalid username"))
				return
			}
			if password != claims["password"] {
				log.Error("Failed to match password: ", err)
				api.RequestErrorHandler(w, errors.New("invalid password"))
				return 
			}
			
			next.ServeHTTP(w, r)
		} else {
			log.Println("Invalid token claims")
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}
	})
}