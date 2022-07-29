package dbconn

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"suiibell/ent"
)

func DBConnection() (*ent.Client, error) {

	errDir := godotenv.Load()
	if errDir != nil {
		log.Panicln("error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// dataSource is the format of postgres connection string
	dataSource := "host=" + dbHost + " port=5432 user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " sslmode=disable"

	// client is the connection pointer of ent
	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client, err
}
