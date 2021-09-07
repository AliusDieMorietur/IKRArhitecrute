package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8765"
	}
	return port
}

func getTimeJSON() string {
	log.Println("time.Now()", time.Now())
	time, err := json.Marshal(map[string]string{"time": time.Now().Format(time.RFC1123)})
	if err != nil {
		log.Fatalln(err)
	}
	return string(time)
}
