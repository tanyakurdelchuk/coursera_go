package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)


func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a name:")
	name, _ := in.ReadString('\n')

	fmt.Println("Please enter an address:")
	address, _ := in.ReadString('\n')

	m := make(map[string]string)
	m["name"] = strings.Trim(name, "\n")
	m["address"] = strings.Trim(address, "\n")

	barr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(string(barr))
}
