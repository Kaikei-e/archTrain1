package dbmanage

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
	"suiibell/ent"
	"suiibell/ent/migrate"
)

func EntInit() {
	p, err := filepath.Abs("../suiibell-backend")
	if err != nil {
		log.Panicln(err)
	}

	errDir := godotenv.Load(filepath.Join(p, ".env"))
	if errDir != nil {
		log.Panicln(errDir)
	}

	errLoad := godotenv.Load(".env")
	if errLoad != nil {
		log.Panicln("error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	//"host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	dataSource := "host=" + dbHost + " port=5432 user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " sslmode=disable"

	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
