package main

import "fmt"

var (
	area = func(a int, b int) int {
		return a * b
	}
)

func main() {
	area := area(10, 20)
	fmt.Println("Area is: ", area)
}