package main

import (
	"fmt"
)

func Extraction(s string) string {
	runes := []rune(s)
	return fmt.Sprintf("%c%c%c%c", runes[0], runes[2], runes[4], runes[6])
}

func main() {
	s := "パタトクカシーー"
	fmt.Println(Extraction(s))
}
