package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string = "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Static files should be served here, but I have no idea on\n"+
			"how to make it work")
	})

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
