package main

import (
	"fmt"
	"knocker/task2/fibbfunction"
)

func main() {
	fmt.Println("welcome")
	defer fmt.Println("fibb numbers printed")
	fibbfunction.Printer(10)
}
