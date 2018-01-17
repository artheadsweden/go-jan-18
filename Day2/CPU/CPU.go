package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		for c := byte('a'); c <= byte('z'); c++ {
			go fmt.Println(string(c))
		}
	}()
	time.Sleep(time.Second * 3)
}
