package main

import "fmt"

var m = map [string]string {
	"c": "Cyan",
	"m": "Magenta",
	"y": "Yellow",
	"k": "Black",
}

func main() {
	// Declaring empty map
	var shopingList = map[string]int{}
	fmt.Println(shopingList)

	// Initializing map
	var people = map[string]int {
		"John": 21,
		"Jane": 22,
		"Jack": 23,
	}
	fmt.Println(people)

	// Map declaration using make function
	var peopleList = make(map[string]int)
	peopleList["John"] = 21
	peopleList["Jane"] = 22
	peopleList["Jack"] = 23
	fmt.Println(peopleList)

	// Accessing items
	fmt.Println(m["c"])

	// Adding items
	m["w"] = "White"
	fmt.Println(m)

	// Updating items
	m["y"] = "lemon yellow"
	fmt.Println(m)

	// Deleting items
	delete(m, "w")
	fmt.Println(m)

	// Iterating over map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Test if an item exists
	c, ok := m["c"]
	fmt.Println(c, ok)
}