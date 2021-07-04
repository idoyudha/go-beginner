package main

import "fmt"

func main() {
	result := factorialRecursive(5)
	fmt.Println(result)
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}