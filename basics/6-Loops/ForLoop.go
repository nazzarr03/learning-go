package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println("Hello World")
		if i == 5 {
			break
		}
		i++
	}

	strings := []string{"Hello", "World", "!"}
	for index, value := range strings {
		fmt.Println(index, value)
	}

	for _, value := range strings {
		fmt.Println(value)
	}

	for index := range strings {
		fmt.Println(index)
	}
}