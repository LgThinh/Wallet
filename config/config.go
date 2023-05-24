package config

import (
	"log"
)

var Config AppConfig

type AppConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func SetEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

}

func LoadEnv() AppConfig {
	return Config

}
