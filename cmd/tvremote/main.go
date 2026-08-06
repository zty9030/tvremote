package main

import (
	"log"
	"net/http"

	"tvremote/internal/server"
	"tvremote/internal/executor"
	"tvremote/internal/input"
	"tvremote/internal/api"
)

func main() {

	adb := executor.NewADBExecutor("host.docke.internal:16416")

	inputService := input.NewService(adb)

	handler := api.NewHandler(inputService)


	log.Println("===================================")
	log.Println(" TV Remote")
	log.Println(" Listening on :8080")
	log.Println("===================================")

	srv := server.New(handler)

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
	}
}
