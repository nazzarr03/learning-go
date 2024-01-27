package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/headers", headers)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", yourHandler()))
}

func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	return
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func yourHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellooooo!")
	})
}