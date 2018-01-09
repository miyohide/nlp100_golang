package main

import "fmt"

func JoinInTurn(s1 string, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var rval string
	for i := 0; i < len(runes1); i++ {
		rval += string(runes1[i]) + string(runes2[i])
	}
	return rval
}

func main() {
	fmt.Println(JoinInTurn("パトカー", "タクシー"))
}
