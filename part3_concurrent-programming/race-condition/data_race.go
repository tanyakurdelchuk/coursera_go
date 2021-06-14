package main

import (
	"fmt"
	"time"
)

/*
A race condition is a phenomenon where the behavior of a program
is dependent on the order in which certain tasks are executed.

A common type of race condition is a data race, which occurs when
two or more goroutines (in the case of Golang) access a variable
and at least one of these accesses updates (writes) the variable.

There is a tool in Golang to detect data races. It can be executed as follows:
go run -race data_race.go
*/

func racecond(cnt *int) {
	*cnt++
	fmt.Println(*cnt)
}

// If you run this program, sometimes it will print 1 2 3 4 5,
// but sometimes it will print the numbers in a different order,
// indicating that its behavior is dependent on the order in which
// the goroutines are executed.
func main() {
	cnt := 0
	go racecond(&cnt)
	go racecond(&cnt)
	go racecond(&cnt)
	go racecond(&cnt)
	go racecond(&cnt)
	time.Sleep(time.Second)
}
