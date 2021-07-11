package fibbfunction

import "fmt"

func fibbfunction() func() int {
	i := 0
	j := 1
	return func() int {
		nextfib := i + j
		i, j = j, nextfib
		return nextfib
	}
}

func Printer(n int) {
	x := fibbfunction()
	for i := 0; i < n; i++ {
		fmt.Println(x())
	}
}
