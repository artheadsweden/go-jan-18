package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

//func (p PairList) Len() int           { return len(p) }
//func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
//func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type String string

func (s String) Split() []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(string(s), f)
}

type WordMap map[string]int

func (w WordMap) Add(s string) {
	w[s]++
}

func (wordFrequence WordMap) rankByWordFrequency() PairList {
	pl := make(PairList, len(wordFrequence))
	i := 0
	for k, v := range wordFrequence {
		pl[i] = Pair{k, v}
		i++
	}
	//sort.Sort(sort.Reverse(pl))
	sort.Slice(pl, func(i, j int) bool { return pl[i].Value > pl[j].Value })
	return pl
}

func main() {
	var s String
	s = "Hi there this is a string that is what it is or is it"
	ss := s.Split()

	counter := make(WordMap)
	for _, word := range ss {
		counter.Add(word)
	}
	pl := counter.rankByWordFrequency()
	fmt.Println(pl)

}
