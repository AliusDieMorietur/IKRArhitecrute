package lib

import (
	"fmt"
	"log"
	"net/http"

	Utils "server/lib/utils"
)

type Routes map[string]func() string
type Handler func(resW http.ResponseWriter, req *http.Request)

func handle(routes Routes) Handler {
	return func(resW http.ResponseWriter, req *http.Request) {
		log.Println("url", req.URL.Path)
		handleRoute := routes[req.URL.Path]
		response := "Not found"
		if handleRoute != nil {
			response = handleRoute()
		}
		fmt.Fprintf(resW, response)
	}
}

func Start(routes Routes) {
	port := Utils.GetPort()
	http.HandleFunc("/", handle(routes))
	log.Printf("Listening at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
