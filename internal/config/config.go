package config

import "os"

type Config struct {
	Port             string
	Env              string
	PostgresUser     string
	PostgresPassword string
}

func GetConfig() Config {
	return Config{
		Port:             os.Getenv("PORT"),
		Env:              os.Getenv("ENV"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
}
