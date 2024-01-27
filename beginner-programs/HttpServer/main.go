package main

import (
	"log"
	"net/http"
	"github.com/nazzarr03/beginner-programs/HttpServer/server"
)

func main() {
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", server.NewServer()))
}