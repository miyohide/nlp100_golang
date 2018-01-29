package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var row1s []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sliced := strings.Split(scanner.Text(), "\t")
		row1s = append(row1s, sliced[0])
	}

	hasVal := make(map[string]bool)
	uniqRow1s := make([]string, 0, len(row1s))

	for _, v := range row1s {
		if _, ok := hasVal[v]; !ok {
			hasVal[v] = true
			uniqRow1s = append(uniqRow1s, v)
		}
	}
	sort.Sort(sort.StringSlice(uniqRow1s))
	for _, v := range uniqRow1s {
		fmt.Println(v)
	}
}
