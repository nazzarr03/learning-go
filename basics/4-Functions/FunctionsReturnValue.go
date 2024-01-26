package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getArea(a int, b int) (area int) {
	area = a * b
	return // Return without specifying variable name.
}

func rectangle(l int, b int) (area int, parameter int) {
	area = l * b
	parameter = 2 * (l + b)
	return
}

// Passing pointer to a function.
func addValue(x *int, y *string) {
	*x = *x + 5
	*y = *y + " World!"
	return
}

func main() {
	sum := add(10, 20)
	fmt.Println("Sum is: ", sum)

	area := getArea(10, 20)
	fmt.Println("Area is: ", area)	

	area, parameter := rectangle(10, 20)
	fmt.Println("Area is: ", area)
	fmt.Println("Parameter is: ", parameter)

	number := 10
	name := "Hello"
	addValue(&number, &name)
	fmt.Println("Number is: ", number)
	fmt.Println("Name is: ", name)
}