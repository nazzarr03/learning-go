package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func main() {
	json_string := `{"name":"Rocky","weight":10}`

	rocky := new(Dog)
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by rocky.
	json.Unmarshal([]byte(json_string), rocky)
	fmt.Println(rocky)

	spot := new(Dog)
	spot.Name = "Spot"
	spot.Weight = 5
	// Marshal returns the JSON encoding of spot.
	jsonStr, _ := json.Marshal(spot)
	fmt.Println(jsonStr)
	fmt.Println(string(jsonStr))	
}