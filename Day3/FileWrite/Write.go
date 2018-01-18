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
	d1 := []byte("Hello there\nThis is a text\nYeah\n")
	fmt.Println(d1)
	err := ioutil.WriteFile("text.txt", d1, 0644)
	check(err)

	f, err := os.Create("data.txt")
	defer f.Close()
	check(err)

	d2 := []byte("Some text - ÅÄÖ")
	n, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes", n)

	n2, err := f.WriteString("Hej hej\n")
	check(err)
	fmt.Printf("Wrote %d bytes", n2)

	f.Sync()

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("This text is buffered\n")
	check(err)
	fmt.Printf("Wrote %d bytes", n3)

	w.Flush()
}
