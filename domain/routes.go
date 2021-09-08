package routes

import (
	"encoding/json"
	"log"
	"time"
)

var ROUTES = map[string]func() string{
	"/":     func() string { return "/ - Help\n/time - Show current time" },
	"/time": getTimeJSON,
}

type Package struct {
	Time string `json:"time"`
}

func getTimeJSON() string {
	time, err := json.Marshal(Package{Time: time.Now().Format(time.RFC1123)})
	if err != nil {
		log.Fatalln(err)
	}
	return string(time)
}
