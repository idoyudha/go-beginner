package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value64)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8) // This will overflow

	var name = "Ido"
	var e = name[0] 
	// Convert byte to string
	var eString string = string(e)

	fmt.Println(name, e, eString)
}