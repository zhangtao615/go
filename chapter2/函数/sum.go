package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	c := sum(1, 2)
	fmt.Printf("sum(1, 2)=%d \n", c)
}
