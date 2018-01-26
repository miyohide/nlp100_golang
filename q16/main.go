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
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	lineSize := len(lines)
	linePerFile := lineSize / *splitNum

	for i := 0; i < *splitNum-1; i++ {
		f, err := os.Create(fmt.Sprintf("%02d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		w := bufio.NewWriter(f)
		for j := linePerFile * i; j < linePerFile*(i+1); j++ {
			w.WriteString(lines[j])
			w.WriteString("\n")
		}
		w.Flush()
	}
	f, err := os.Create(fmt.Sprintf("%02d.txt", *splitNum-1))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := linePerFile * (*splitNum - 1); i < linePerFile*(*splitNum); i++ {
		w.WriteString(lines[i])
		w.WriteString("\n")
	}
	w.Flush()
}
