package main

import "fmt"

func main() {
	var ido Customer
	ido.Name = "Ido"
	ido.Address = "Indonesia"
	ido.Age = 25

	fmt.Println(ido)
	fmt.Println(ido.Name, ido.Address, ido.Name)

	naruto := Customer {
		Name:    "Naruto",
		Address: "Konoha",
		Age:     30,
	}

	fmt.Println(naruto)

	// Call struct method
	naruto.sayHi("Sasuke") 
}

type Customer struct {
	Name, Address string
	Age           int
}

// Struct function or struct method
func (customer Customer) sayHi(name string) { 
	fmt.Println("Hello", name, "my name is", customer.Name)
}
