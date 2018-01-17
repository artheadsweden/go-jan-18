package main

import "fmt"

func a(i int) {
	fmt.Println(i)
	i++
}

func b(i *int) {
	fmt.Println(*i)
	*i++
}

func main() {
	x := 10
	// var xptr *int = &x
	xptr := &x
	fmt.Println(xptr, *xptr)

	*xptr++
	fmt.Println(x, *xptr)

	a(x)
	fmt.Println("x=", x)
	// a(xptr) ERROR
	a(*xptr)
	fmt.Println(x)

	b(&x)
	fmt.Println(x)
	b(xptr)
	fmt.Println(x)

	iptr := new(int)
	fmt.Println(*iptr)

	var xptrptr **int = &xptr
	var xptrptrptr ***int = &xptrptr

	fmt.Println(**xptrptr)
	fmt.Println(***xptrptrptr)

	***xptrptrptr = 24
	fmt.Println(x)
	b(**xptrptrptr)
}
