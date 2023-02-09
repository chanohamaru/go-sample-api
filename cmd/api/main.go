package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()
}
