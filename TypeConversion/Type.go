package main

import "fmt"

func main() {
	f := 3.44
	//b := true
	//	var i int = int(b) ERROR
	var i int = int(f)
	fmt.Println(i)
	fmt.Printf("f = %T \n", f)
}
