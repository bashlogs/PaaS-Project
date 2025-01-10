package main

import (
	"database/sql"
	"fmt"
	"log"
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