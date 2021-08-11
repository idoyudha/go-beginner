package main

import "fmt"

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	// passing data into variadic function
	numbers := []int{10, 20, 30}
	total = sumAll(numbers...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}