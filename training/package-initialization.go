package main

import (
	"database"
	_ "hash"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}