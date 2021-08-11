package main

import "fmt"

func main() {
	var name string
	name = "Ido"
	fmt.Println(name)

	var phone = "Samsung"
	fmt.Println(phone)

	var age = 30 // type int
	fmt.Println(age)

	var num int8 = 10 // type int
	fmt.Println(num)

	fullName := "Ido Yudhatama" // first declaration
	fmt.Println(fullName)

	fullName = "Ido Widya"
	fmt.Println(fullName)

	var (
		firsName = "Ido Widya"
		lastName = "Yudhatama"
	)

	fmt.Println(firsName, lastName)

}
