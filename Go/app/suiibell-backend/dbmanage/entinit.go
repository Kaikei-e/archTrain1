package dbmanage

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"suiibell/ent"
	"suiibell/ent/migrate"
)

func EntInit() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
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
