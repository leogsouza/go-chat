package utils

import "os"

func DefaultPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	return port
}
