package main

import "fmt"

func main() {
	a := 5
	b := 10

	func() {
		area := a * b
		fmt.Println("Area is: ", area)
	} ()
}