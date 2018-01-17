package main

import "fmt"

type Cat struct {
	Sound string
}

func (c Cat) Speak() {
	fmt.Println(c.Sound)
}

type Dog struct {
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

type Animal interface {
	Speak()
}

func main() {
	d1 := Dog{"Bow wow"}
	d2 := Dog{"Wiff wiff"}
	c1 := Cat{"Miew"}
	c2 := Cat{"Yawn"}

	animals := []Animal{d1, c1, d2, c2}
	for _, a := range animals {
		a.Speak()
	}
}
