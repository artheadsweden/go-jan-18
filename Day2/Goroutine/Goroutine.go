package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
	}
}

func main() {
	//f("Sekvens")
	go f("Goroutine")

	go func(name string) {
		for i := 0; i < 3; i++ {
			fmt.Println(name, i)
		}
	}("Anon Goroutine")
	time.Sleep(time.Second * 2)
}
