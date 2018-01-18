package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ops uint64 = 0

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				//atomic.AddUint64(&ops, 1)
				ops++
				time.Sleep(time.Millisecond * 1)
			}
		}()
	}

	wg.Wait()
	finalOps := ops //atomic.LoadUint64(&ops)
	fmt.Println(finalOps)

}
