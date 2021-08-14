package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	/**
	PushBack: add in the end
	PushFront: add in the front
	*/
	data.PushBack("Jakarta")
	data.PushBack("Bandung")
	data.PushBack("Surabaya")
	data.PushFront("Jayapura")

	fmt.Println(data.Front().Prev())
	fmt.Println(data.Front().Next())

	// Print first data
	fmt.Println(data.Front().Value)
	// Print last data 
	fmt.Println(data.Back().Value)

	// Iterate through double linked list start from front to end
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Iterate through double linked list start from end to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}