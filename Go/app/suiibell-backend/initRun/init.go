package initRun

import (
	"github.com/joho/godotenv"
	"suiibell/dbmanage"
)

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load the .env file")
	}

	dbmanage.EntInit()

	return nil

}
