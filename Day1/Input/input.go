package main

import "fmt"

func main() {
	fmt.Print("String1: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("String2: ")
	var second string
	fmt.Scanln(&second)
	fmt.Println(first, second)
}
