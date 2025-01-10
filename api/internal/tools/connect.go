package tools

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface struct {
	DB *sql.DB
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

		dataSourceName := "user=postgres password=khadde dbname=test sslmode=disable"

		db, dbErr := sql.Open("pgx", dataSourceName)
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
