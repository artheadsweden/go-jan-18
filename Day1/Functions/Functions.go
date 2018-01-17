package main

import "fmt"

func myFuncA(a int, b string) int {
	return a
}

func myFuncB(a, b, c int, d string) {

}

func myFuncC(a int) (int, bool) {
	return 0, false
}

func myFuncD() (x int, y bool) {
	x = 3
	y = true
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func runner(f func(int) int, x int) {
	f(x)
}

func main() {
	fmt.Println(myFuncA(1, "Hej"))

	x, y := myFuncC(2)
	fmt.Println(x, y)

	a, b := myFuncD()
	fmt.Println(a, b)

	i, j := swap(3, 4)
	fmt.Println(i, j)

	runner(
		func(a int) int {
			fmt.Println(a)
			return a
		}, 44)

	myFunc := func(a int) int {
		fmt.Println(a)
		return a
	}
	runner(myFunc, 44)

	func(a int) {
		fmt.Println(a)
	}(4)
}
