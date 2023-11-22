package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ServerConfig struct {
	HTTPHost string
	HTTPPort string
	DBHost   string
	DBPort   string
	DBName   string
	DB       string
	DBUser   string
	DBPwd    string
}

func NewConfig() ServerConfig {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	serverConfig := ServerConfig{
		HTTPHost: os.Getenv("HTTPHOST"),
		HTTPPort: os.Getenv("HTTPPORT"),
		DBHost:   os.Getenv("DBHOST"),
		DBPort:   os.Getenv("DBPORT"),
		DBName:   os.Getenv("DBNAME"),
		DB:       os.Getenv("DB"),
		DBUser:   os.Getenv("DBUSER"),
		DBPwd:    os.Getenv("DBPWD"),
	}
	return serverConfig
}

func NewLinks() []string {
	return []string{
		os.Getenv("AGIFYLINK"),
		os.Getenv("GENDERIZELINK"),
		os.Getenv("NATINALIZELINK"),
	}
}
