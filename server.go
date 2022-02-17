package main

import (
	"fmt"
	"log"
	"net/http"

	"dogtrainer/handlers"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/dogs", handlers.DogsHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
