package main

import "fmt"

func first() {
	fmt.Println("First")
}	

func second() {
	fmt.Println("Second")
}

func main() {
	// Defer is used to delay the execution of a function until the surrounding function returns.
	defer second()
	first()
}