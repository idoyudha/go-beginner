package main

import (
	"errors"
	"fmt"
)

func main() {
	var expError error = errors.New("Get Error")
	fmt.Println(expError.Error())

	// result, err := Divider(100, 5)
	result, err := Divider(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}

func Divider(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider can't be 0")
	} else {
		result := value / divider 
		return result, nil
	}
}