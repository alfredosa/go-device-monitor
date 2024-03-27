package main

import (
	"log"
	"os"
)

func getJWTSecretKey() []byte {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(secretKey)
}
