package config

import (
	"os"
	"strconv"
)

type APIConfig struct {
	API_BASE_URL string
	API_PORT     int
}

type DBConfig struct {
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}
type Config struct {
	API APIConfig
	DB  DBConfig
}

func NewConfig() *Config {
	return &Config{
		API: APIConfig{
			API_BASE_URL: getEnv("API_BASE_URL", ""),
			API_PORT:     getEnvAsInt("API_PORT", 8080),
		},
		DB: DBConfig{

			DB_HOST:     getEnv("DB_HOST", ""),
			DB_PORT:     getEnvAsInt("DB_PORT", 5432),
			DB_USER:     getEnv("DB_USER", ""),
			DB_PASSWORD: getEnv("DB_PASSWORD", ""),
			DB_NAME:     getEnv("DB_NAME", ""),
		},
	}

}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if valueInt, err := strconv.Atoi(valueStr); err == nil {
			return valueInt
		}
		return defaultValue
	}
	return defaultValue
}
