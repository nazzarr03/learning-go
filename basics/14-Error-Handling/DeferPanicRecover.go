package main

import "fmt"

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}

	fmt.Println("End: recoveryFunction")
}

func executePanic() {
	defer recoveryFunction()
	panic("Panic")
	fmt.Println("End: executePanic")
}

func main() {
	executePanic()
	fmt.Println("End: main")
}