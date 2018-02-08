package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
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
	// 正規表現はここを参考にバッククォーテーションで
	// 囲んでいる。https://ashitani.jp/golangtips/tips_regexp.html
	reg := regexp.MustCompile(`File:([^|]+)`)
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
				if m := reg.FindString(line); len(m) > 0 {
					fmt.Printf("ファイル名: %s\n", m)
				}
			}
		}
	}
}
