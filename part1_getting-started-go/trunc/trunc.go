package main

import "fmt"

func main() {
	var num float64
	fmt.Println("Please type a floating point number:")
	fmt.Scan(&num)
	fmt.Println(int(num))
}
