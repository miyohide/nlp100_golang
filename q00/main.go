package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	reverses := make([]rune, len(runes))

	for i := 0; i < len(runes); i++ {
		reverses[i] = runes[len(runes)-i-1]
	}
	return string(reverses)
}

func main() {
	s := "stressed"
	fmt.Println(Reverse(s))
}
