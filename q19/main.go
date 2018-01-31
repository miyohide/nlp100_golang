package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	line string
	val  int
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

	lines := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "\t")
		_, ok := lines[splited[0]]
		if ok {
			lines[splited[0]] = lines[splited[0]] + 1
		} else {
			lines[splited[0]] = 1
		}
	}

	sorted := List{}
	for k, v := range lines {
		e := Entry{k, v}
		sorted = append(sorted, e)
	}
	sort.Sort(sorted)

	for i := range sorted {
		fmt.Printf("%s, %d\n", sorted[i].line, sorted[i].val)
	}

}
