package main

import "fmt"

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("num is 1")
	case 2:
		fmt.Println("num is 2")
	case 3:
		fmt.Println("num is 3")	
	default:
		fmt.Println("num is not 1, 2 or 3")
	}

	dayOfWeek := 3
	switch dayOfWeek{
	case 1:
		fmt.Println("Go to work")
		fallthrough // this will execute the next case
	case 2:
		fmt.Println("Buy some bread")
		fallthrough
	case 3:
		fmt.Println("Visit a doctor")
		fallthrough
	case 4:
		fmt.Println("Buy some food")
		fallthrough
	case 5:
		fmt.Println("See your family")
	default:
		fmt.Println("No information available for that day")
	}

	switch number := 5; {
	case number < 5:
		fmt.Println("Smaller than 5")
	case number == 5:
		fmt.Println("Five")
	case number > 5:
		fmt.Println("Bigger than five")
	default:
		fmt.Println("No information about the number")
	}
}