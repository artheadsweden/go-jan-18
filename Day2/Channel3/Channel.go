package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for n := range ch {
		fmt.Println(n)
	}

	ch2 := make(chan int, 4)
	go fibonacci(cap(ch2), ch2)

	for value, ok := <-ch2; ok; value, ok = <-ch2 {
		fmt.Println(value)
	}

}
