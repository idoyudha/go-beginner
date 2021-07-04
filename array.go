package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ido"
	names[1] = "Yudhatama"

	fmt.Println(names[0], names[1])

	var values = [3]int {
		90,80,85,
	}

	fmt.Println(values)

	values[1] = 95 // replace array values
	fmt.Println("Length of array =>", len(values))
	fmt.Println("Array after replacement =>", values)
}