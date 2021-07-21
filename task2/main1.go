package main

import (
	"fmt"

	"github.com/egor200201/golanghomework/task2/fibbfunction"
)

func main() {
	fmt.Println("welcome")
	defer fmt.Println("fibb numbers printed")
	fibbfunction.Printer(10)
}
