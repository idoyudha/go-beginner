package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ido",
		"address": "Jember",
	}

	// Add data 
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}