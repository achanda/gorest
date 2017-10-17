package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: os.Getenv("USER"),
			Password: os.Getenv("PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("HOST"),
		},
	}
}
