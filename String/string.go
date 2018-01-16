package main

import "fmt"

func main() {
	message := "\"Who are you?\" he asked"

	fmt.Println(message)
	message2 := `Multi line
string`
	fmt.Println(message2)

	message3 := "√ \u221A \U0000221a"
	fmt.Println(message3)

	char1 := "A"
	fmt.Println(len(char1))

	char2 := "√"
	fmt.Println(len(char2))

	str1 := "abc"
	str2 := "åäö"
	fmt.Println(len(str1))
	fmt.Println(len(str2))

	str3 := "漢字"
	fmt.Println(len(str3))

	fmt.Println(str1[:2])
	fmt.Println(str2[:2])
	fmt.Println(str3[:2])

	s1 := "\u00c5"       // Å
	s2 := "\u212B"       // Ångström symbolen Å
	s3 := "\u0041\u030a" // A med ring över

	fmt.Println(s1, s2, s3)

	fmt.Println(s1 == s2, s1 == s3, s2 == s3)

}
