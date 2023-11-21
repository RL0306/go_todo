package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetValueFromEnvFile(key string) string {
	return os.Getenv(key)
}

func InitialiseLoggingConfig(path string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading in environment file")
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening loading file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Log file created")
}
