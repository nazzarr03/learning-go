package main

import (
	"fmt"
	"strconv"
	"reflect"
)
func main() {
	// Converting a string to an float64
	s := "3.141592653589"
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f)
	
	// Converting a string to an boolean
	str := "true"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)

	// Converting a float to an string
	var flo float64 = 3.141592653589
	var strFlo string = strconv.FormatFloat(flo, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(strFlo))

	// Converting a float to an int

	var f32 float32 = 3.141592653589
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
}