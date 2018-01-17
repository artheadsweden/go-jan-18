package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	Sq := Square{10}
	Ci := Circle{10}
	PrintArea(Sq)
	PrintArea(Ci)
}
