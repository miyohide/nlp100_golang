package main

import "fmt"

func Template(x int, y string, z float64) string {
	return fmt.Sprintf("%d時の%sは%2.1f", x, y, z)
}

func main() {
	fmt.Println(Template(12, "気温", 22.4))
}
