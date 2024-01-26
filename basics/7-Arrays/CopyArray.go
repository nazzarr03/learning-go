package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	y := x // Copying the array

	z := &x // Copying the pointer

	fmt.Println(x)
	fmt.Println(y)

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)
}