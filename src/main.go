package main

import (
	"log"
	"net/http"
	"os"
	"whiteboard/handler"
)

func main() {
	http.HandleFunc("/getUpdates", handler.GetUpdates)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	http.ListenAndServe(":"+port, nil)
}
