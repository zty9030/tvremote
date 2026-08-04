package main

import (
	"log"
	"net/http"

	"tvremote/internal/server"
)

func main() {

	log.Println("===================================")
	log.Println(" TV Remote")
	log.Println(" Listening on :8080")
	log.Println("===================================")

	srv := server.New()

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
	}
}
