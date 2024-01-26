package main

import "fmt"

func main() {
	hello("GO")
	add(10, 20)
}

func hello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a int, b int) {
	fmt.Println("Sum is: ", a + b)
}