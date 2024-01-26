package main

import "fmt"

func main() {
	var array [5]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array[0], array[1])

	x := [5]int{1, 2, 3, 4, 5}
	var y [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(y)

	k := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(k))

	a := [5]int{1: 1, 3: 25}
	fmt.Println(a)
}