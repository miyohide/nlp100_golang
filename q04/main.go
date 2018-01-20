package main

import (
	"fmt"
	"strings"
)

func WordIndex(s string) map[string]int {
	s = strings.Replace(s, ".", "", -1)
	words := strings.Split(s, " ")

	rval := make(map[string]int)

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		switch i + 1 {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			rval[string(runes[0])] = i + 1
		default:
			rval[string(runes[0])+string(runes[1])] = i + 1
		}
	}
	return rval
}

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	fmt.Println(WordIndex(s))
}
