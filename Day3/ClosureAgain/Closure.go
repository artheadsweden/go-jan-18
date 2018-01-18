package main

import "fmt"

func main() {
	x := 10
	f := func() func() {
		x := x
		return func() {
			fmt.Println(x)
		}
	}()
	x++
	f()
}
