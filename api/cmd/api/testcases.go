package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bashlogs/PaaS_Project/api/internal/tools"
	"github.com/golang-jwt/jwt"
)

func Print() {
    fmt.Println("Hello World")
}

func CreateTable(){
	// connStr := "user=postgres password=khadde dbname=test sslmode=disable" // for pq driver
	connStr := "postgres://postgres:khadde@localhost:5432/test?sslmode=disable" // for pqx driver
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	_, err = db.Exec(`Create Table users (
		user_id integer NOT NULL, 
		name varchar not null, 
		age integer not null, 
		PRIMARY KEY(user_id))`)

	if err != nil {
		log.Fatal("Failed to create database: ", err)
	}

	fmt.Println("Database created successfully")
}

func Query(username string, pwd string){
	database, err := tools.ConnectToDatabase()
    if err != nil {
        log.Fatal("Database connection error: ", err)
        return
    }

	var password string
	err = database.DB.QueryRow("select password from users where username=$1", username).Scan(&password)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("user not found")
		}
		log.Fatal("database error: ", err)
	}

	if password == pwd {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}

	// if pwd {
	// 	fmt.Println("Login successful")
	// } else {
	// 	fmt.Println("Login failed")
	// }

}

var secretKey = []byte("khadde")

func JWT_Authenticate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err // Return the specific error
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verify token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return false, errors.New("token has expired")
			}
		} else {
			return false, errors.New("invalid expiration in token")
		}

		return true, nil
	}

	return false, errors.New("invalid token")
}