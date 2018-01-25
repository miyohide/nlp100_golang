package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	splitNum = flag.Int("n", 10, "ファイルの分割数")
	fileName = flag.String("f", "", "読み込みファイル名")
)

func main() {
	flag.Parse()

	fp, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(fp)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	lineSize := len(lines)
	linePerFile := lineSize / *splitNum

	for i := 0; i < *splitNum; i++ {
		f, err := os.Create(fmt.Sprintf("%02d.txt", i))
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)
		for j := 0; j < linePerFile; j++ {
			w.WriteString(lines[(*splitNum)*i+j])
			w.WriteString("\n")
		}
		w.Flush()
	}
}
