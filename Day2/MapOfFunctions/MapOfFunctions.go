package main

import (
	"fmt"
)

func main() {
	calc := make(map[string]func(int, int) int)

	calc["add"] = func(a, b int) int { return a + b }
	calc["sub"] = func(a, b int) int { return a - b }
	calc["mul"] = func(a, b int) int { return a * b }
	calc["div"] = func(a, b int) int { return a / b }

	for name, function := range calc {
		fmt.Println(name, function(20, 4))
	}
}
