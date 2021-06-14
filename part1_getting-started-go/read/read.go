package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the FULL PATH of the file you wish to read:")
	fname, err := in.ReadString('\n')
	file, err := os.Open(strings.Trim(fname, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n := strings.Fields(scanner.Text())
		names = append(names, Name{fname: n[0], lname: n[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range(names) {
		fmt.Printf("first name: %s\t last name: %s\n", name.fname, name.lname)
	}
}
