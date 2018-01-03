package main

import (
	"bufio"
	"os"
)

func main() {
	filename := os.Args[1]

	rfp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer rfp.Close()

	wfp1, werr1 := os.Create("col1.txt")
	if werr1 != nil {
		panic(err)
	}
	defer wfp1.Close()

	wfp2, werr2 := os.Create("col2.txt")
	if werr2 != nil {
		panic(err)
	}
	defer wfp2.Close()

	scanner := bufio.NewScanner(rfp)

	line := 1
	for scanner.Scan() {
		if line%2 == 0 {
			wfp1.Write(([]byte)(scanner.Text() + "\n"))
		} else {
			wfp2.Write(([]byte)(scanner.Text() + "\n"))
		}
		line++
	}
}
