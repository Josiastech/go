package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Return connection to the server with a string of text
// Utilizes r.URL.Path[1:] to append connection URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

// Direct all traffic to port :8080 utilizing default handler
// Ther handler manages the traffic flowing into port :8080
// throughout the API's defined functions or routes
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}

// Struct are typed collection of fields. They're useful for
// grouping data together to form records.
type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to GO LANG the API"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	w.Write(b)
}
