package main

import "fmt"

func main() {
	sayHello()
	sayWithParameter("Ido", "Yudhatama")

	fmt.Println("------")

	fmt.Println(getHello(""))
	fmt.Println(getHello("Ido"))
	
	fmt.Println("------")

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println("------")

	// if we pass second returning value
	first, _ := getFullName()
	fmt.Println(first)

	fmt.Println(getData())
}

func sayHello() {
	fmt.Println("Hello!")
}

// function with parameter
func sayWithParameter(first string, last string) {
	fmt.Println("Hello", first, last)
}

// function with retuning a value
func getHello(name string) string {
	if name == "" {
		return "Hello"
	} else {
		return "Hello " + name
	}
}

// Function that returning multiple value
func getFullName() (string, string) {
	return "Ido", "Yudhatama"
}

// named return values 
func getData() (name string, university string, age int) {
	name = "Ido"
	university = "ITS"
	age = 25
	return
}
