package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument: ", args)
	// go run os.go ido
	// Argument:  [C:\Users\yudha\AppData\Local\Temp\go-build1920849922\b001\exe\os.exe ido]

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	// Setting config
	username := os.Getenv("GO_USERNAME")
	password := os.Getenv("GO_PASSWORD")
	fmt.Println(username, password)
}