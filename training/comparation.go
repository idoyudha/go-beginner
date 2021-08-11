package main

import "fmt"

func main() {
	var name1 = "Ido"
	var name2 = "Ido"

	var result bool = name1 == name2
	fmt.Println("name1 = name2,",result)

	var value1 = 100
	var value2 = 200 

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
}