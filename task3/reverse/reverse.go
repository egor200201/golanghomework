package main

import "fmt"

func rev(elements []int64) []int64 {
	reversed := []int64{}

	for i := range elements {
		n := elements[len(elements)-1-i]

		reversed = append(reversed, n)
	}

	return reversed
}
func main() {
	elements := []int64{1, 2, 5, 15}
	fmt.Println("reverse oder")
	fmt.Println("Before : ", elements)
	fmt.Println("After : ", rev(elements))
}
