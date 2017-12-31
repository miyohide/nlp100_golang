package main

import "fmt"

func main() {
	s1 := []rune("パトカー")
	s2 := []rune("タクシー")

	for i := range s1 {
		fmt.Printf("%c%c", s1[i], s2[i])
	}
	fmt.Printf("\n")
}
