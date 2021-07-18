package main

import (
	"fmt"
	"sort"
)

func printSorted(m map[int]string) {

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(m[k])
	}
}
func main() {
	var m = map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(m)
}
