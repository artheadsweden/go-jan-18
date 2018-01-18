package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("text.txt")
	check(err)
	fmt.Print(string(data))

	f, err := os.Open("data.txt")
	check(err)
	b1 := make([]byte, 14)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("Read %d bytes, got %s\n", n1, string(b1))

	o1, err := f.Seek(12, os.SEEK_SET)
	check(err)
	fmt.Println(o1)

	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("Read %d bytes, got %s\n", n2, string(b2))

	r1 := bufio.NewReader(f)
	s1, err := r1.ReadString('\n')
	check(err)
	fmt.Println(s1)
	f.Close()

	f1, err := os.Open("data.txt")
	r2 := bufio.NewReader(f1)
	check(err)

	for s2, err := r2.ReadString('\n'); err == nil; s2, err = r2.ReadString('\n') {
		fmt.Print(s2)
	}
	check(err)

	f1.Close()
}
