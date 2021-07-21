package main

import "fmt"

func max(str []string) string {
	maxvalue := ""
	for _, value := range str {
		if len(value) > len(maxvalue) {
			maxvalue = value
		}
	}
	return maxvalue
}
func main() {
	str := []string{"one", "two", "three"}
	fmt.Println(" the longest word from the slice of strings (the first if there are more than one)")
	fmt.Print(max(str))
}
