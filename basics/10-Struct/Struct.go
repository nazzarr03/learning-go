package main

import "fmt"

type Animal struct {
	name   string
	weight int
}

func main() {
	var dog Animal
	dog.name = "Dog"
	dog.weight = 10
	fmt.Println(dog)

	var cat Animal = Animal{
		name:   "Cat",
		weight: 5,
	}
	fmt.Println(cat)

	var bird = new(Animal)
	bird.name = "Bird"
	bird.weight = 1
	fmt.Println(bird)

	var monkey = &Animal{
		name:   "Monkey",
		weight: 20,
	}
	fmt.Println(monkey)

	fmt.Println(bird==monkey)

	monkey2 := monkey
	monkey2.name = "Monkey2"
	fmt.Println(monkey2)
}