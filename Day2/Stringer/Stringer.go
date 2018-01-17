package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main() {
	p1 := Person{"Athur Dent", 42}
	p2 := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(p1)
	fmt.Println(p2)
}
