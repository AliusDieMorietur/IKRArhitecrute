package main

import (
	"fmt"
	"log"
	"net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {
	log.Println("url", r.URL.Path)
	f := END_POINTS[r.URL.Path]
	respose := "Not found"
	if f != nil {
		respose = f()
	}
	fmt.Fprintf(w, respose)
}

func main() {
	port := getPort()
	http.HandleFunc("/", serve)
	log.Printf("Listening at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
