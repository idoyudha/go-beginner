package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.49))
	fmt.Println(math.Floor(1.80))
	fmt.Println(math.Round(2.80))
	fmt.Println(math.Max(2, 7))
	fmt.Println(math.Min(2, 7))
}