package config

import (
	"os"
)

func GetDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://user:password@localhost:5432/shipping_db?sslmode=disable"
	}
	return url
}

func GetApplicationPort() string {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "50053"
	}
	return port
}
