package main

import "fmt"

func Cipher(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] = 219 - runes[i]
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(Cipher("hoge123"))
}
