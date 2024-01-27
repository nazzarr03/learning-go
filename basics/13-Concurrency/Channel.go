package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func querySite(url string, requestBody chan int) {
	defer wg.Done()
	
	fmt.Println("HTTP Request: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Read Body: ", url)
	// ioutil.ReadAll reads the entire response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body length: ", len(body))

	// <- is the "receive" operator when used with channels
	requestBody <- len(body)
}

func main() {
	// Declaring a unbuffered channel
	requestBody := make(chan int)

	wg.Add(1)
 
	// Writing to the channel
	go querySite("https://github.com/nazzarr03", requestBody)

	// Reading from the channel
	fmt.Println(<-requestBody)
	
	wg.Wait()

	// Closing the channel
	close(requestBody)
}