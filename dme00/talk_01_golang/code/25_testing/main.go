package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1, 2))
}

// Sum adds a to b and returns the result.
func Sum(a int, b int) int {
	return a + b
}
