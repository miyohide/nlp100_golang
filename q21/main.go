package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type PageStruct struct {
	Title, Text string
}

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var p PageStruct
	reader := bufio.NewReaderSize(r, 4096)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		dec := json.NewDecoder(strings.NewReader(line))
		if err = dec.Decode(&p); err != nil {
			panic(err)
		}
		if p.Title == "イギリス" {
			lines := strings.Split(p.Text, "\n")
			for _, line := range lines {
				if strings.Contains(line, "Category:") {
					fmt.Printf("%v\n", line)
				}
			}
		}
	}
}
