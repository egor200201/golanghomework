package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("average of array:", Average(a))
}

func Average(a [6]int) float64 {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return float64(sum) / float64(len(a))
}
