package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func DogsHandler(w http.ResponseWriter, r *http.Request) {
	// console output
	fmt.Println("inside DogHandler")

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// response
		msg := []byte("Ajax, Buddy, Champ")
		_, err := w.Write(msg)

		// error handling
		if err != nil {
			log.Fatal(err)
		} else {
			// log output
			log.Println("responded with a list of dogs")
		}

	default:
		msg := "Method not allowed"
		http.Error(w, msg, http.StatusMethodNotAllowed)
		log.Println(msg)
	}
}
