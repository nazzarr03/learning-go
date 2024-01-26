package main

import "fmt"

func main() {
	x := true
	if x {
		fmt.Println("true")
	}

	y := 1
	if y > 0 {
		fmt.Println("y is greater than 0")
	} else {
		fmt.Println("y is less than 0")
	}

	grade := 5
	if grade == 1 {
		fmt.Println("You have an A")
	} else if grade > 1 && grade < 5 {
		fmt.Println("You have no A but you are positive")
	} else {
		fmt.Println("You grade is negative")
	}

	if a := 10; a == 10 {
		fmt.Println(a)
	}
}