package main

import "fmt"

func main() {

	elements := [...]int64{1, 2, 5, 15}

	fmt.Println("Before : ", elements)

	reversed := []int64{}

	for i := range elements {
		n := elements[len(elements)-1-i]

		reversed = append(reversed, n)
	}

	fmt.Println("After : ", reversed)

}
