package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animalsActionsMap := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print("> ")
		var animal, action string
		fmt.Scan(&animal, &action)

		switch action {
		case "eat":
			animalsActionsMap[animal].Eat()
		case "move":
			animalsActionsMap[animal].Move()
		case "speak":
			animalsActionsMap[animal].Speak()
		}
	}
}
