package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			x++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			x++
		}
	}()

	wg.Wait()
	fmt.Println(x)
}
