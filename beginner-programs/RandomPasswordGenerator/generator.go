package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	// UnixNano returns a random number. We use this as the seed for the random number generator.
	rand.Seed(time.Now().UnixNano())

	numbers := "0123456789"
	specialCharacters := "!@#$%^&*()_+{}[]:;?/.,"
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" + 
				  specialCharacters + numbers

	password := make([]byte, length)

	for i := 0; i < length; i++{
		password[i] = characters[rand.Intn(len(characters))]
	}
	
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func main() {
	pass:= generatePassword(10)
	fmt.Println(pass)
}