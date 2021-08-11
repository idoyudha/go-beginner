package main

import "fmt"

func main() {
	// var data int = Hai(1)
	var data interface{} = Hai(1)
	fmt.Println(data)
}

func Hai(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "None"
	}
}