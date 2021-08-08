package main

import "fmt"

func main() {
	var ido Person 
	ido.Name = "Ido"

	SayHello(ido)

	tiger := Animal {
		Name: "Samsung",
	}

	SayHello(tiger)
}

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}


type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}


type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}