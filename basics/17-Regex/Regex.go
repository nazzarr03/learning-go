package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "[0-9]+"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
	}

	// Test if the pattern works
	str := "Some text 1 2 3 "
	if re.MatchString(str) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	// Return match
	result := re.FindString(str)
	fmt.Println(result)

	// Return multiple matches
	results := re.FindAllString(str, -1)
	fmt.Println(results)

	// Replace matches
	replaced := re.ReplaceAllString(str, "x")
	fmt.Println(replaced)

	// Submatches
	str1 := "Hello @world@ Match this"
	sub_re, _ := regexp.Compile("@(.*)@")
	m := sub_re.FindStringSubmatch(str1)
	if len(m) > 1 {
		fmt.Println(m[1])
	}
}