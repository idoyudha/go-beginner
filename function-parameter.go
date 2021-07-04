package main

import "fmt"

func main() {
	sayHelloWithFilter("Ido", spamFilter)
	sayHelloWithFilter("Kaka", spamFilter)
}

// using type declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func spamFilter(name string) string {
	if name == "Ido" {
		return "..."
	} else {
		return name
	}
}