package dbconn

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DBConnection() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	//dsn := "host=" + dbHost + "user=" + dbUser + "password=" + dbPass + "dbname=" + dbName + "port=3306 sslmode=disable"
	dataSource := "host=" + dbHost + " port=5432 user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to initialize db connection")
	}

	return db, err
}
