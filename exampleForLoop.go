package main

import (
	"fmt"
)

func main() {
	for i, j, k := 0, 10, 100; i < 100 && j > 0  && k < 1000; i, j, k = i+1, j-1, k+2 {
		fmt.Println("Value of i, j, k : [", i, j, k," ]")
	}
}
