package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	for sum := 0; ; {
		sum++
		if sum > 30 {
			break
		}
	}
	fmt.Println(sum)
}
