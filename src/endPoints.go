package main

import (
	"encoding/json"
	"log"
	"time"
)

var END_POINTS = map[string]func() string{
	"/":     func() string { return "/ - Help\n/time - Show current time" },
	"/time": getTimeJSON,
}

func getTimeJSON() string {
	time, err := json.Marshal(Package{Time: time.Now().Format(time.RFC1123)})
	if err != nil {
		log.Fatalln(err)
	}
	return string(time)
}
