package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename1 := os.Args[1]
	filename2 := os.Args[2]

	fp1, err := os.Open(filename1)
	if err != nil {
		panic(err)
	}
	defer fp1.Close()

	fp2, err := os.Open(filename2)
	if err != nil {
		panic(err)
	}
	defer fp2.Close()

	file1_lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(fp1)
	for scanner.Scan() {
		file1_lines = append(file1_lines, scanner.Text())
	}
	if scanner_err := scanner.Err(); scanner_err != nil {
		panic(scanner_err)
	}

	file2_lines := make([]string, 0, 100)
	scanner = bufio.NewScanner(fp2)
	for scanner.Scan() {
		file2_lines = append(file2_lines, scanner.Text())
	}
	if scanner_err := scanner.Err(); scanner_err != nil {
		panic(scanner_err)
	}

	for i := 0; i < len(file1_lines); i++ {
		fmt.Printf("%s\t%s\n", file1_lines[i], file2_lines[i])
	}
}
