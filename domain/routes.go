package routes

import (
	"encoding/json"
	"log"
	"time"
)

type Package struct {
	Time string `json:"time"`
}

var ROUTES = map[string]func() string{
	"/":     RouteRoot,
	"/time": RouteTime,
}

func RouteRoot() string { return "/ - Help\n/time - Show current time" }
func RouteTime() string { return getTimeJSON() }

func getTimeJSON() string {
	time, err := json.Marshal(Package{Time: time.Now().Format(time.RFC3339)})
	if err != nil {
		log.Fatalln(err)
	}
	return string(time)
}
