/*
	Slice Internals

	-------------------------------------
    |   addr    |     3    |      5      |
	|  Pointer  |  Length  |   Capacity  |
    -------------------------------------
       |
       V
	--------------------------
	| 10 | 20 | 30 | 0  | 0  |
	--------------------------
	   0    1   2    3    4
*/

package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[:]
	arr1[1] = 99

	fmt.Println(slice1)
	slice1 = append(slice1, 6)
	fmt.Println(slice1)
	arr1[2] = 45
	fmt.Println(slice1)

	slice2 := arr1[1:4]
	fmt.Println(slice2)
	fmt.Println(arr1)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, 77)
	fmt.Println(arr1)

	slice3 := []int{1, 2, 3}
	slice3 = append(slice3, 4)
	fmt.Println(slice3)

	slice4 := make([]int, 4, 5) // Type, Len, Cap
	fmt.Println(slice4)
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))

	for i, value := range slice4 {
		fmt.Println(i, value)
	}
}
