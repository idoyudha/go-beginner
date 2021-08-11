package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ido Yudhatama", "Ido"))
	fmt.Println(strings.Split("Ido Yudhatama", " "))
	fmt.Println(strings.ToLower("Ido Yudhatama"))
	fmt.Println(strings.ToUpper("Ido Yudhatama"))
	fmt.Println(strings.Trim("  Ido Yudhatama  ", " "))
	fmt.Println(strings.ReplaceAll("do do do do", "do", "did"))
}