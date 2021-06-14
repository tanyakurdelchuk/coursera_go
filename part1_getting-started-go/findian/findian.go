package main

import (
	"fmt"
	"strings"
)

func findian(str string) string {
	s := strings.ToLower(str)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		return "Found!"
	} else {
		return "Not Found!"
	}
}


func main() {
	fmt.Println("Please enter a string:")
	var str string
	fmt.Scan(&str)
	s := strings.ToLower(str)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
