package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type server struct {
	HTTPHost string
	HTTPPort string
}

type db struct {
	DBHost string
	DBPort string
	DBName string
	DB     string
	DBUser string
	DBPwd  string
}

func NewDbConfig() *db {
	return &db{
		DBHost: os.Getenv("DBHOST"),
		DBPort: os.Getenv("DBPORT"),
		DBName: os.Getenv("DBNAME"),
		DB:     os.Getenv("DB"),
		DBUser: os.Getenv("DBUSER"),
		DBPwd:  os.Getenv("DBPWD"),
	}
}

func NewServerConfig() server {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	s := server{
		HTTPHost: os.Getenv("HTTPHOST"),
		HTTPPort: os.Getenv("HTTPPORT"),
	}
	return s
}

func NewLinks() []string {
	return []string{
		os.Getenv("AGIFYLINK"),
		os.Getenv("GENDERIZELINK"),
		os.Getenv("NATINALIZELINK"),
	}
}
