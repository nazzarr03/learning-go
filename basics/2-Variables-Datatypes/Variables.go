package main

import "fmt"

func main() {
	var i int
	var s string

	i = 5
	s = "Hello World!"

	fmt.Println(i)
	fmt.Println(s)

	var k  int = 35
	fmt.Println(k)

	j := 10
	fmt.Println(j)

	firstName, lastName := "Nazar", "Arik"
	fmt.Println(firstName + " " + lastName)

	var (
		name = "Nazar"
		age = 20
	)
	fmt.Println(name)
	fmt.Println(age)
}