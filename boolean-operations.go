package main

import "fmt"

func main() {
	var value = 90
	var abs = 80

	var passValue = value > 80
	var passAbs = abs >= 80

	var pass = passValue && passAbs
	fmt.Println(pass)
}