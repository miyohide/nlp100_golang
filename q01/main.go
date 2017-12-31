package main

import (
	"fmt"

	"golang.org/x/exp/utf8string"
)

func main() {
	s := "パタトクカシーー"
	s2 := utf8string.NewString(s)
	fmt.Println(s2.Slice(0, 1) + s2.Slice(2, 3) + s2.Slice(4, 5) + s2.Slice(6, 7))
}
