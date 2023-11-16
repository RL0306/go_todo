package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvironmentFile() error {
	return godotenv.Load()
}

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}

func GetValueFromEnvFile(key string) string {
	return os.Getenv(key)
}
