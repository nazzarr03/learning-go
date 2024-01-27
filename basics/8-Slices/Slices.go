package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x []int
	fmt.Println(reflect.ValueOf(x).Kind())

	// Creating a slice using make function
	var y = make([]string, 10, 20)
	fmt.Println(len(y), cap(y))

	// Inıtialize the slice with values using a slice literal
	var z = []int{1, 2, 3, 4, 5}
	fmt.Println(len(z), cap(z))
	fmt.Println(z)

	// Creating a slice using the new keybord
	var a = new([50]int)[0:10]
	fmt.Println(len(a), cap(a))
	fmt.Println(a)

	// Add items using append function
	var b = make([]int, 1, 10)
	fmt.Println(b)
	b = append(b, 20)
	fmt.Println(b)

	// Access slice items
	var c = []int{1, 2, 3, 4, 5}
	fmt.Println(c[0])
	fmt.Println(c[1:3])

	// Change item values
	var d = []int{1, 2, 3, 4, 5}
	fmt.Println(d)
	d[0] = 10
	fmt.Println(d)

	// Copy slice into another slice
	var e = []int{1, 2, 3, 4, 5}
	var f = []int{6, 7, 8, 9, 10}
	copy(f, e)
	fmt.Println(f)

	// Append a slice to an existing one
	var g = []int{1, 2, 3, 4, 5}
	var h = []int{6, 7, 8, 9, 10}
	g = append(g, h...)
	fmt.Println(g)
}