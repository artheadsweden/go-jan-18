package main

import "fmt"

func main() {
	println("Hej hej")
	print("Hej igen\n")
	fmt.Println("Hej världen")

	a, b, c := "Hi", "there", "folks"
	fmt.Println(a, b, c)
	fmt.Print(a, b, c, "\n")

	fmt.Printf("Int repr: %v - %d \n", 10, 10)
	fmt.Printf("Float repr: %v - %f %.2f \n", 10.3, 10.3, 10.3)
	fmt.Printf("Type of: %T \n", false)
}
