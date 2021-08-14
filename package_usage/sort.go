package main

import (
	"fmt"
	"sort"
)

func main() {
	users := []User{
		{"Ido", 25},
		{"Doi", 22},
		{"Ari", 23},
		{"Ira", 26},
	}
	fmt.Println(users) // before sort
	sort.Sort(UserSlice(users))
	fmt.Println(users) // after sort
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}