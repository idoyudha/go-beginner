package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run flag.go -host=localhost -user=root -password=root
	// First is argument, 2nd is default value, and third is description
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()
	// First is pointer then value -> *
	fmt.Println("Host: ", host, *host)
	fmt.Println("User: ", user, *user)
	fmt.Println("Password: ", password, *password)
}