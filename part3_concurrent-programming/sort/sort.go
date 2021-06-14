package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortIntegers(part []int, c chan []int) {
	fmt.Printf("Sorting part: %v\n", part)
	sort.Ints(part)
	c <- part
}

func merge(lo, hi []int) []int {
	result := make([]int, len(lo)+len(hi))
	i := 0
	for len(lo) > 0 && len(hi) > 0 {
		if lo[0] < hi[0] {
			result[i] = lo[0]
			lo = lo[1:]
		} else {
			result[i] = hi[0]
			hi = hi[1:]
		}
		i++
	}

	for j := 0; j < len(lo); j++ {
		result[i] = lo[j]
		i++
	}
	for j := 0; j < len(hi); j++ {
		result[i] = hi[j]
		i++
	}

	return result
}

func main() {
	fmt.Println("Please enter a sequence of integers:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var numbers []int
	for _, n := range strings.Fields(scanner.Text()) {
		ni, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		numbers = append(numbers, ni)
	}

	const numParts = 4
	sizeParts := len(numbers) / numParts
	c := make(chan []int, numParts)

	go sortIntegers(numbers[:sizeParts], c)
	go sortIntegers(numbers[sizeParts:2*sizeParts], c)
	go sortIntegers(numbers[2*sizeParts:3*sizeParts], c)
	go sortIntegers(numbers[3*sizeParts:], c)

	parts := make([][]int, numParts)
	for i := 0; i < numParts; i++ {
		parts[i] = <-c
	}

	firstHalf := merge(parts[0], parts[1])
	secondHalf := merge(parts[2], parts[3])
	sortedParts := merge(firstHalf, secondHalf)
	fmt.Printf("Final result: %v", sortedParts)
}
