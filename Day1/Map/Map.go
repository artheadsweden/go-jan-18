package main

import "fmt"

func main() {
	// nil map
	var persons1 map[string]int
	fmt.Println(persons1)
	//persons1["Paul"] = 10

	persons2 := make(map[string]int)
	fmt.Println(persons2)
	persons2["Paul"] = 10
	fmt.Println(persons2)
	persons2["Anna"] = 20
	fmt.Println(persons2)
	age := persons2["Anna"]
	fmt.Println(age)

	age = persons2["Peter"]
	fmt.Println(age)
	fmt.Println(persons2)
	var ok bool
	if age, ok = persons2["Peter"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("Key not found")
	}

	for key, value := range persons2 {
		fmt.Println(key, value)
	}

	delete(persons2, "Peter")
	fmt.Println(persons2)
}
