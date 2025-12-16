package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return get("ENV")
}

func GetDataSourceURL() string {
	return get("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	p := get("APPLICATION_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatalf("invalid port: %s", p)
	}
	return port
}

func GetPaymentServiceUrl() string {
	return get("PAYMENT_SERVICE_URL")
}


func get(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing environment variable: %s", key)
	}
	return v
}
