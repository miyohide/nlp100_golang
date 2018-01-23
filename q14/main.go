package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	headLines = flag.Int("n", 10, "表示する行数")
	fileName  = flag.String("f", "", "読み込みファイル名")
)

func main() {
	flag.Parse()

	fp, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(fp)
	i := 1

	for sc.Scan() {
		if i > *headLines {
			break
		} else {
			i++
			fmt.Println(sc.Text())
		}
	}
}
