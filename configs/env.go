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

	pDBUser := getEnvVar("POSTGRES_USER")
	pDBPassword := getEnvVar("POSTGRES_PASSWORD")
	pDBName := getEnvVar("POSTGRES_DB")

	pDBUrl := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		pDBUser, pDBPassword, pDBName,
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
