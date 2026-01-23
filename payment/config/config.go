package config

import (
	"log"
	"os"
)

func get(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing env var %s", key)
	}
	return val
}

func GetDatabaseURL() string     { return get("DATABASE_URL") }
func GetApplicationPort() string { return get("APPLICATION_PORT") }
func GetEnvironment() string     { return get("ENV") }
