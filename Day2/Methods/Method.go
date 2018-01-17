package main

import "fmt"

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func Area(s Square) float64 {
	return s.Side * s.Side
}

//================================

type MyFloat float64

func (f MyFloat) Squared() MyFloat {
	return f * f
}

//================================

type Point struct {
	X, Y int
}

func (p Point) MoveX(step int) {
	p.X += step
}

func (p *Point) MoveY(step int) {
	p.Y += step
}

func main() {
	s1 := Square{8}
	fmt.Println(Area(s1))
	fmt.Println(s1.Area())

	var f MyFloat = 3.16
	fmt.Println(f.Squared())

	p1 := Point{10, 20}
	p1.MoveX(3)
	p1.MoveY(3)
	fmt.Println(p1)
	p2 := &p1
	p2.MoveY(3)
	fmt.Println(p1)
}
