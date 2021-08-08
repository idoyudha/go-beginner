package main

import "fmt"

func main() {
	ido := Man{"Ido"}
	ido.Married()
	fmt.Println(ido.Name)
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}