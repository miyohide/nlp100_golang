package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	tailLines = flag.Int("n", 10, "表示する末尾の行数")
	fileName  = flag.String("f", "", "読み込みファイル名")
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
	for i := lineSize - *tailLines; i < lineSize; i++ {
		fmt.Println(lines[i])
	}
}
