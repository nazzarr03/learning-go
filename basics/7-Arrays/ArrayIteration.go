package main

import "fmt"

func main() {	
	x := [5]int{1, 2, 3, 4, 5}

	// Basic for loop	
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Getting the index and value
	for index, value := range x {
		fmt.Println(index, value)
	}

	// Getting only the value
	for _, value := range x {
		fmt.Println(value)
	}

	// Getting only the index
	for index := range x {
		fmt.Println(index)
	}

	// Range and counter
	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}
}