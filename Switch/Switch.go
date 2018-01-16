package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("This program is running on: ")

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("MS Windows")
	default:
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("MS Windows")
	default:
		fmt.Println(os)
	}

	x := 12
	// ...

	switch {
	case x < 10, x == 12:
		fmt.Println("Small x")
	case x < 100:
		fmt.Println("Medium x")
	default:
		fmt.Println("Large x")
	}

	switch x := 10; {
	case x >= 0 && x < 5:
		fmt.Println("< 5")
	case x >= 0 && x < 10:
		fmt.Println("< 10")
	case x == 10:
		fmt.Println("== 10")
		fallthrough
	default:
		fmt.Println("Bad value for x")
	}

	switch {
	case f():
		if g() {
			goto nextCase
		}
		h()
		break
	nextCase:
		fallthrough
	default:
		error()
	}
}

func f() bool {
	return true
}

func g() bool {
	return true
}

func h() {
	fmt.Println("All is good")
}

func error() {
	fmt.Println("All is bad")
}
