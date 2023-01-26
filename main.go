package main

import (
	"fmt"
	"log"
	"net/http"

	"go-server/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":" + port, r))
}
