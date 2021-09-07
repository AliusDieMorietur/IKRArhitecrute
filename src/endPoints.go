package main

var END_POINTS = map[string]func() string {
	"/": func () string {return "/ - Help\n/time - Show current time"},
	"/time": getTimeJSON,
}