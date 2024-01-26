package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func partialMultiply(a int) func(int) int {
	return func(b int) int {
		return multiply(a, b)
	}
}

func main() {
	multiply := partialMultiply(10)
	fmt.Println(multiply(20))
}