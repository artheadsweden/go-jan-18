package main

import "fmt"

const pi = 3.14

const (
	startValue = 23
	endValue   = 44
)

const (
	first = iota
	second
)

const (
	third = iota
	forth
)

const (
	_          = iota
	kb float64 = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

func main() {
	fmt.Println(first, second, third, forth)
	fmt.Println(kb, mb, gb, tb, pb, eb, zb, yb)
}
