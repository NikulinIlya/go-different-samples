package main

import (
	"fmt"
)

// Calculate : add 2
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing Sample")
	result := Calculate(2)
	fmt.Println(result)
}
