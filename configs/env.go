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

	pDBUrl := os.Getenv("POSTGREDB_URL")
	if pDBUrl == "" {
		log.Fatal("You must set your 'POSTGREDB_URL' environmental variable.")
	}

	return pDBUrl
}
