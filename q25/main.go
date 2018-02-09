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
	template_start_reg := regexp.MustCompile(`^\\{\\{基礎情報 .*$`)
	template_end_reg := regexp.MustCompile(`^\\}\\}$`)
	name_reg := regexp.MustCompile(`^\\||=\\s.*$`)
	value_reg := regexp.MustCompile(`^\\|.* = `)
	is_template_reg := regexp.MustCompile(`^\\|.* = .*$`)

	dict := make(map[string]string)
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
			template_start := false
			template_end := false
			for _, line := range lines {
				if template_end {
					break
				}
				if template_start {
					isTemplate := is_template_reg.MatchString(line)
					template_end = template_end_reg.MatchString(line)

					if isTemplate {
						name := name_reg.ReplaceAllString(line, "")
						value := value_reg.ReplaceAllString(line, "")
						dict[name] = value
					}
				} else {
					template_start = template_start_reg.MatchString(line)
				}
			}
		}
	}

	for name, value := range dict {
		fmt.Printf("%v = %v\n", name, value)
	}
}
