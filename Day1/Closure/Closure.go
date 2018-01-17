package main

import (
	"fmt"
	"math"
)

func outer(name string) func(int) int {
	inner := func(i int) int {
		fmt.Println(name)
		return i
	}
	return inner
}

type floatFunc func(float64) float64

func funcFact(p float64) floatFunc {
	return func(b float64) float64 {
		return math.Pow(b, p)
	}
}

func main() {
	f := outer("Anna")
	fmt.Println(f(3))
	fmt.Println(f(4))
	f2 := outer("Peter")
	fmt.Println(f2(3))
	fmt.Println(f(4))

	square := funcFact(2)
	cube := funcFact(3)

	fmt.Println(square(3))
	fmt.Println(square(4))

	fmt.Println(cube(3))
	fmt.Println(cube(4))

	sq := func(p float64) func(float64) float64 {
		return func(b float64) float64 {
			return math.Pow(b, p)
		}
	}(2)
	fmt.Println(sq(3))

	ff := func(p float64) func(float64) float64 {
		return func(b float64) float64 {
			return math.Pow(b, p)
		}
	}
	sq2 := ff(2)
	fmt.Println(sq2(3))
}
