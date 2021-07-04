package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Looping", counter)
		counter++
	}

	// For with statement 
	for count := 1; count <= 10; count++ {
		fmt.Println("Looping 2", count)
	}

	// Looping in slice 
	slice := []string{"ITS", "ITB", "UGM"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Looping with range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	// Looping in map
	person := make(map[string]string)
	person["name"] = "Ido"
	person["university"] = "ITS"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}