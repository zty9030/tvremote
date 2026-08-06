package main

import (
	"log"
	"net/http"
	"os"
	"tvremote/internal/server"
	"tvremote/internal/executor"
	"tvremote/internal/input"
	"tvremote/internal/api"
)

func main() {

    var exec executor.AndroidExecutor

    if addr := os.Getenv("ADB_ADDRESS"); addr != "" {
        exec = executor.NewADBExecutor(addr)
    } else {
        exec = executor.NewShellExecutor()
    }

	inputService := input.NewService(exec)

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
