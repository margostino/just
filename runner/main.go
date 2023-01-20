package main

import (
	handler "github.com/margostino/just/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Jobs)
	log.Println("Starting Just (dev) Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
