package main

import (
	"log"
	"net/http"

	"github.com/PregOfficial/word-game/handler"
)

func main() {
	http.Handle("/", ServeFrontend())

	http.HandleFunc("/api/start", handler.StartGame)
	http.HandleFunc("/api/guess", handler.Guess)

	log.Println("Server running at 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
