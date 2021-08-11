package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert input to boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// Convert input to int
	number, err := strconv.ParseInt("10000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// Convert input to binary (2)
	value := strconv.FormatInt(1000, 2)
	fmt.Println(value)
}