package main

import (
	"fmt"
	"strings"
)

func main() {
	// Creating a string
	str := "Hello world!"
	sentence := "This is, a sentence."
	numbers := []string{"one", "two", "three"}
	path := "/Desktop/test"

	// Convert string to lowercase
	lower := strings.ToLower(str)
	fmt.Println(lower)

	// Convert string to uppercase
	upper := strings.ToUpper(str)
	fmt.Println(upper)

	// Check if string contains another string
	if strings.Contains(str, "world") {
		fmt.Println("Yes this string exists")
	}

	// Get/Print character range from string
	fmt.Println(str[0:5])
	fmt.Println(str[6:])

	// Split a string
	split := strings.Split(sentence, ",")
	fmt.Println(split[0])
	fmt.Println(split[1])

	// Split a string by whitespace
	fields := strings.Fields(sentence)
	fmt.Println(fields)

	// Join an array of strings
	joinStr := strings.Join(numbers, ",")
	fmt.Println(joinStr)

	// Replace a string (Takes a number of how many replacements it should do -1 = all of them)
	newStr := strings.Replace(str, "world", "earth", -1)
	fmt.Println(newStr)

	// HasPrefix
	isAbsolute := strings.HasPrefix(path, "/")
	fmt.Println(isAbsolute)

	// HasSuffix
	isTextFile := strings.HasSuffix(path, "/")
	fmt.Println(isTextFile)

	// Count characters in string
	count := strings.Count(str, "o")
	fmt.Println(count)

	// Determine string length
	length := len(str)
	fmt.Println(length)

	// Repeat a string
	repeat := strings.Repeat(str, 2)
	fmt.Println(repeat)

	// Trim whitespace from string
	trim := strings.TrimSpace(str)
	fmt.Println(trim)
}