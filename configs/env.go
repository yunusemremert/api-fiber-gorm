package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvPostgreDBUrl() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	pDBHost := getEnvVar("POSTGRES_HOST")
	pDBUser := getEnvVar("POSTGRES_USER")
	pDBPassword := getEnvVar("POSTGRES_PASSWORD")
	pDBName := getEnvVar("POSTGRES_DB")
	pDBPort := getEnvVar("POSTGRES_PORT")

	pDBUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pDBHost, pDBUser, pDBPassword, pDBName, pDBPort,
	)

	return pDBUrl
}

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("You must set your '%s' environmental variable.", key)
	}

	return value
}
