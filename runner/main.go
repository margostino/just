package main

import (
	handler "github.com/margostino/earth-station-api/api/nasa"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Nasa)
	log.Println("Starting Lumos Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
