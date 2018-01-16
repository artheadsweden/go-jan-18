package main

import "fmt"

func sum(numbers ...int) (total int) {
	//fmt.Printf("%T \n", numbers)
	for _, value := range numbers {
		total += value
	}
	return
}

func pln(head string, tail ...string) {
	fmt.Print(head)

	if len(tail) > 0 {
		fmt.Print(" ")
		head, tail = tail[0], tail[1:]
		pln(head, tail...)
	} else {
		fmt.Println()
	}
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(10, 20, 30, 40, 50, 60))
	fmt.Println(sum())

	pln("Hej")              // Hej \n
	pln("Hej", "på", "dig") // Hej på dig \n
}
