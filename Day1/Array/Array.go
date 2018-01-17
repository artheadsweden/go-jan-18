package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr1 [5]int
	fmt.Println(arr1)
	arr1[1] = 4
	arr1[2] = 5
	fmt.Println(arr1)

	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr3)

	arr4 := [5]int{1: 10, 3: 15}
	fmt.Println(arr4)

	ptrArr := [5]*int{1: new(int), 3: new(int)}

	*ptrArr[1] = 12
	fmt.Println(ptrArr)

	arr5 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(arr5)

	arr6 := [4][2]int{1: {20, 21}, 3: {30, 31}}
	arr6[0][0] = 10
	arr6[0][1] = 11
	fmt.Println(arr6)

	var bigArray [1e6]int
	fmt.Println(unsafe.Sizeof(bigArray))
	foo2(&bigArray)
}

func foo1(x [1e6]int) {

}

func foo2(x *[1e6]int) {

}
