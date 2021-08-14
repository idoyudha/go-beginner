package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())

	// parse, _ := time.Parse(time.RFC3339, "2021-8-12")
	// fmt.Println(parse)
}