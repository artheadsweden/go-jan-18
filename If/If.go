package main

import "fmt"

func main() {
	x := 10

	// ...

	if x < 15 {
		fmt.Println("x < 15")
	} else {
		fmt.Println("x >= 15")
	}

	computedValue := func(x int) int { return x * x }

	if computedValue(10) <= 110 {
		fmt.Println("Comuted value is", computedValue(10))
	}

	value := computedValue(10)
	if value <= 110 {
		fmt.Println("Comuted value is", value)
	}

	if value := computedValue(10); value <= 110 {
		fmt.Println("Comuted value is", value)
	}

}
