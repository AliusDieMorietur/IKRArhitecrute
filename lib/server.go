package lib

import (
	"fmt"
	"log"
	"net/http"

	Utils "server/lib/utils"
)

type Routes map[string]func() string

func handle(routes Routes) func(resW http.ResponseWriter, req *http.Request) {
	return func(resW http.ResponseWriter, req *http.Request) {
		log.Println("url", req.URL.Path)
		handleRoute := routes[req.URL.Path]
		respose := "Not found"
		if handleRoute != nil {
			respose = handleRoute()
		}
		fmt.Fprintf(resW, respose)
	}
}

func Start(routes Routes) {
	port := Utils.GetPort()
	http.HandleFunc("/", handle(routes))
	log.Printf("Listening at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
