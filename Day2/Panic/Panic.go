package main

import (
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatalln("Fel", e)
	}

}

func main() {
	data, err := ioutil.ReadFile("/x/apa.txt")
	check(err)
}
