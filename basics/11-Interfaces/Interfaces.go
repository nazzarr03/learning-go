package main

import "fmt"

type User interface {
	PrintName(name string)
}

type Vehicle interface {
	Alert() string
}

type Usr int

type Car struct{}

func (usr Usr) PrintName(name string) {
	fmt.Println(name)
	fmt.Println(usr)
}

func (c Car) Alert() string {
	return "Alert!"
}

// Define an empty interface - Often used for functions that accept any type 
func Print(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func main() {
	var user1 User
	user1 = Usr(1)
	user1.PrintName("John")

	c := Car{}
	fmt.Println(c.Alert())

	Print("Hello", 1, 2, 3, 4, 5)
}