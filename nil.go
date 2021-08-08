package main

import "fmt"

func main() {
	var person map[string]string = NewMap("")
	if person == nil {
		fmt.Println("Zero data")
	} else {
		fmt.Println(person)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}