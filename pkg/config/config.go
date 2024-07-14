package config

import (
	"os"
	"strconv"
)

type Config struct {
	MaxGoroutines int
}

func LoadConfig() Config {
	return Config{
		MaxGoroutines: getEnvAsInt("MAX_GOROUTINES", 10),
	}
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
