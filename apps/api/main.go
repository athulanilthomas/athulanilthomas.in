package main

import (
	"log"
	"net/http"
	"os"

	"github.com/athulanilthomas/www/api/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router, err := router.BuildRouter()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
