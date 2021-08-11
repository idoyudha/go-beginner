package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"Febuary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Append
	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "December")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Ido"
	newSlice[1] = "Yudhatama"

	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// Copy slice 
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Array vs Slice 
	array := [3]int{1,2,3} // array
	slice := []int{1,2,3} // slice
	fmt.Println(array, slice)
}