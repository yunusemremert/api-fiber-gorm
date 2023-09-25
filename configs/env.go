package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvPostgreDBUrl() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	postgreDBUrl := os.Getenv("POSTGREDB_URL")
	if postgreDBUrl == "" {
		log.Fatal("You must set your 'POSTGREDB_URL' environmental variable.")
	}

	return postgreDBUrl
}
