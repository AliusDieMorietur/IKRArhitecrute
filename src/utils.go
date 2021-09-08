package main

import (
	"os"
)

type Package struct {
	Time string `json:"time"`
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8765"
	}
	return port
}
