package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	s := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind."
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		length := len(runes)
		if length > 4 {
			new_runes := make([]rune, length)
			indexes := rand.Perm(length - 2)
			new_runes[0] = runes[0]
			new_runes[length-1] = runes[length-1]

			for i := range indexes {
				new_runes[i+1] = runes[indexes[i]+1]
			}
			words[i] = string(new_runes)
		}
	}
	fmt.Println(strings.Join(words, " "))
}
