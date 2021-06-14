package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	var numbers []int
	fmt.Println("Please enter a sequence of integers:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, n := range strings.Fields(scanner.Text()) {
		ni, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		numbers = append(numbers, ni)
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}
