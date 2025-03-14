package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

var (
	ErrUnauthorized = errors.New("unauthorized access")
	ErrInvalidToken = errors.New("invalid token")
)

var SecretKey = []byte("khadde")

// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("authToken")
// 		if err != nil {
// 			if errors.Is(err, http.ErrNoCookie) {
// 				http.Error(w, "Authorization token missing", http.StatusUnauthorized)
// 				return
// 			}
// 			log.Error("Error retrieving cookie:", err)
// 			http.Error(w, "Invalid cookie", http.StatusBadRequest)
// 			return
// 		}

// 		tokenString := cookie.Value
// 		if tokenString == "" {
// 			http.Error(w, "Authorization token is empty", http.StatusUnauthorized)
// 			return
// 		}

// 		// Parse and validate the token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return secretKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			log.Error("Invalid token:", err)
// 			api.RequestErrorHandler(w, ErrInvalidToken)
// 			return
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			log.Error("Failed to parse token claims")
// 			api.RequestErrorHandler(w, ErrUnauthorized)
// 			return
// 		}

// 		// Validate token expiration
// 		if exp, ok := claims["exp"].(float64); ok {
// 			if time.Now().Unix() > int64(exp) {
// 				log.Error("Token has expired")
// 				api.RequestErrorHandler(w, errors.New("token has expired"))
// 				return
// 			}
// 		} else {
// 			log.Error("Invalid expiration in token")
// 			api.RequestErrorHandler(w, errors.New("invalid token expiration"))
// 			return
// 		}

// 		// Validate email and password against the database
// 		email, ok := claims["email"].(string)
// 		if !ok {
// 			log.Error("Invalid email in token")
// 			api.RequestErrorHandler(w, errors.New("invalid email in token"))
// 			return
// 		}

// 		password := claims["password"]
// 		database, err := tools.ConnectToDatabase()
// 		if err != nil {
// 			log.Error("Database connection error:", err)
// 			api.InternalErrorHandler(w)
// 			return
// 		}

// 		var storedPassword string
// 		err = database.DB.QueryRow("SELECT password FROM users WHERE email=$1", email).Scan(&storedPassword)
// 		if err != nil {
// 			log.Error("Email not found in database:", err)
// 			api.RequestErrorHandler(w, errors.New("invalid email"))
// 			return
// 		}

// 		if password != storedPassword {
// 			log.Error("Password mismatch")
// 			api.RequestErrorHandler(w, errors.New("invalid password"))
// 			return
// 		}

// 		log.Info("User authorized:", email)
// 		next.ServeHTTP(w, r)
// 	})
// }

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authToken")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "Authorization token missing", http.StatusUnauthorized)
				return
			}
			log.Error("Error retrieving cookie:", err)
			http.Error(w, "Invalid cookie", http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value
		if tokenString == "" {
			http.Error(w, "Authorization token is empty", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			log.Error("Invalid token:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}
		}

		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Invalid email in token", http.StatusUnauthorized)
			return
		}

		// Add email to request context
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}