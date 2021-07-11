package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome")
	defer fmt.Println("fibb numbers")
	fibbfunction.Printer(10)
}
