package tools

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5" // PostgreSQL driver
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface struct {
	DB *sql.DB
}

func (d *DatabaseInterface) Close() {
	panic("unimplemented")
}

var (
	dbInstance *DatabaseInterface
	once       sync.Once
)

// ConnectToDatabase initializes a singleton database connection
func ConnectToDatabase() (*DatabaseInterface, error) {
	var err error
	once.Do(func() {
		// dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		// 	os.Getenv("DB_USER"),     // Use environment variables
		// 	os.Getenv("DB_PASSWORD"),
		// 	os.Getenv("DB_NAME"),
		// )
		
		err := godotenv.Load(".env")
        if err != nil {
            log.Errorf("Error loading .env file: %v", err)
            return
        }

		// dataSourceName := "user=postgres password=khadde dbname=test sslmode=disable"
		dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME"),
        )

		db, dbErr := sql.Open("postgres", dataSourceName)
		if dbErr != nil {
			err = dbErr
			log.Errorf("Failed to open database: %v", err)
			return
		}

		// Verify the connection
		if pingErr := db.Ping(); pingErr != nil {
			err = pingErr
			log.Errorf("Failed to ping database: %v", err)
			return
		}

		log.Info("Successfully connected to the database")
		dbInstance = &DatabaseInterface{DB: db}
	})

	return dbInstance, err
}
