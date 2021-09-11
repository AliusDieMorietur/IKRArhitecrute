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
	"/":     rootHandler,
	"/time": getTimeJSON,
}

func rootHandler() string { return "/ - Help\n/time - Show current time" }
func getTimeJSON() string {
	time, err := json.Marshal(Package{Time: time.Now().Format(time.RFC1123)})
	if err != nil {
		log.Fatalln(err)
	}
	return string(time)
}
