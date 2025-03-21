package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang API!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Golang server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}