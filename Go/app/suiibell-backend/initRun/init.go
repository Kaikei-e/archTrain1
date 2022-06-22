package initRun

import "github.com/joho/godotenv"

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("failed to load the .env file")
	}

}
