package main

import (
	"fmt"
	"reflect"
)

func main() {
	printMultipleStrings("Hello", "World", "!")
	printMultipleVariables(1, "green", false, 3.14, []string{"a", "b", "c"})
}

// Passing multiple attributes using a variadic function.
func printMultipleStrings(strs ...string) {
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
}

// Pass multiple different datatypes.
func printMultipleVariables(vars ...interface{}) {
	for _, v := range vars{
		fmt.Println(v, "-->", reflect.ValueOf(v).Kind())
	}
}