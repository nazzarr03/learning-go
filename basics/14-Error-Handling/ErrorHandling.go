package main

import (
	"errors"
	"fmt"
	"math"
)

type Error interface {
	Error() string
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative number")
	} else {
		z := 1.0
		min_delta := 0.000000001
		delta := 1.0

		i := 0
		for ; math.Abs(delta) > min_delta; i++ {
			delta = (z*z - x) / (2 * z)
			z -= delta
		}

		fmt.Println("Iterations:", i)
		return z, nil
	}
}

func main() {
	err := errors.New("Error message")
	fmt.Println(err)

	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}