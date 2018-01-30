package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	line string
	val  float64
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].val == l[j].val {
		return (i < j)
	}
	return (l[i].val > l[j].val)
}

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// キーをファイルの1行、値を第3カラムのものにしたmap
	lines := make(map[string]float64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sliced := strings.Split(line, "\t")
		val, _ := strconv.ParseFloat(sliced[2], 64)
		lines[line] = val
	}

	sorted := List{}
	for k, v := range lines {
		e := Entry{k, v}
		sorted = append(sorted, e)
	}
	sort.Sort(sorted)

	for i := range sorted {
		fmt.Println(sorted[i].line)
	}
}
