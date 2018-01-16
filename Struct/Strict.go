package main

import "fmt"

type Point struct {
	X, Y int
	Name string
}

func main() {
	var p Point
	fmt.Println(p)
	p.X = 10
	p.Y = 11
	p.Name = "Punkt A"
	fmt.Println(p)

	p2 := Point{13, 14, "Punkt B"}
	fmt.Println(p2)

	p3 := Point{Y: 14, Name: "Origo"}
	fmt.Println(p3)
}
