package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s %d\n", words[i], len(words[i]))
	}
}
