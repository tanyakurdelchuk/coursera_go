package main

import (
	"fmt"
	"sort"
	"strconv"
)


func main() {
	var sli []int
	var n string
	for true {
		fmt.Println("Please enter a number (type X to quit):")
		fmt.Scan(&n)
		if n == "X" {
			break
		}
		ni, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Wrong type of input (expected int). Ignoring input...")
			continue
		}

		sli = append(sli, ni)
		sort.Ints(sli[:])
		fmt.Println(sli)
	}
}
