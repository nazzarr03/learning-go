package main

import "fmt"

func main() {
	// Basic while loop
	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}

	// Do while loop
	num := 0
	for {
		fmt.Println(num)
		if num == 5 {
			break
		}
		num++
	}
}