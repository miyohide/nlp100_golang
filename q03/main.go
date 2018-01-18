package main

import (
	"fmt"
	"strings"
)

func CountWords(s string) ([]string, []int) {
	wordlens := []int{}
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)

	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		wordlens = append(wordlens, len(words[i]))
	}
	return words, wordlens
}

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	words, lens := CountWords(s)

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s %d\n", words[i], lens[i])
	}
}
