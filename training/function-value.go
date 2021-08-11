package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye // function pass to variable
	result := sayGoodBye("Ido")
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Goodbye " + name
}