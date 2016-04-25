package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This webserver is running. You are currently in %q.", html.EscapeString(r.URL.Path))
	})

	fmt.Println("Webserver running on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}