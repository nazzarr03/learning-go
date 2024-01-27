package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func querySite(url string) {
	fmt.Println("HTTP Request: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Read Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Body length: ", len(body))
}

func main() {
	// Declaring a unbuffered channel
	go querySite("https://github.com/nazzarr03")
	go querySite("https://google.com")

	// Anonymous function
	go func(x int) {
		fmt.Println("Number: ", x)
	} (42)

	time.Sleep(5 * time.Second)
}