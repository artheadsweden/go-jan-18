package main

import "fmt"

type User struct {
	Fname, Ename string
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I don't know how to handle type %T \n", v)
	}
}

func main() {
	var i interface{}
	fmt.Printf("%v, %T \n", i, i)
	i = 3.17
	fmt.Printf("%v, %T \n", i, i)
	i = "Hi there"
	fmt.Printf("%v, %T \n", i, i)
	i = 17
	fmt.Printf("%v, %T \n", i, i)
	i = User{"Anna", "Jones"}
	fmt.Printf("%v, %T \n", i, i)

	i = 10
	do(i)
	i = "Hi"
	do(i)
	i = 4.56
	do(i)

	i = 10

	if s, ok := i.(string); ok {
		fmt.Println("i is a string", s)
	} else {
		fmt.Println("i is not a string", s)
	}
}
