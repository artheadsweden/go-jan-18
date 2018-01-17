package main

import (
	"fmt"
	"sync"
)

type SafeX struct {
	x   int
	mux sync.Mutex
}

func (sx *SafeX) inc() {
	sx.mux.Lock()
	sx.x++
	sx.mux.Unlock()
}

func main() {
	x := SafeX{x: 0}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			x.inc()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			x.inc()
		}
	}()

	wg.Wait()
	fmt.Println(x.x)
}
