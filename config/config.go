package config

import (
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
}

func InitConf() Config {
	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "admin"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "admin123"),
		DB_NAME:        GetOrDefault("DB_NAME", "fits_health"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "database-1.cv8cnjvlnjwz.us-west-1.rds.amazonaws.com"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "AlphaWolf"),
	}
}

func GetOrDefault(envKey, defaultValue string) string {
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}

	return defaultValue
}
