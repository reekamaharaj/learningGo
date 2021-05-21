package main

import (
	"fmt"
)

type AnimalMethods interface {
	Eat()
	Move()
	Speak()
	Name() string
}

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Name() {
	fmt.Println("\n Response:", a.name, "\n")
}

func (a *Animal) Eat() {
	fmt.Println("\n Response:", a.food, "\n")
}

func (a *Animal) Move() {
	fmt.Println("\n Response:", a.locomotion, "\n")
}

func (a *Animal) Speak() {
	fmt.Println("\n Response:", a.noise, "\n")
}

func main() {

}

// Animal interface, struct with three string fields, three types cow, snake, bird
// User can create animals of one of the three types
// User can get information about an animal they created
// Animal has a name and a type
// Animal interface has three methods, methods take no arguments and return no values

// Accept one response at a time
// Print out a response
// New prompt with a new line
// Program loops forever

// User commands "new animal" or "query"

// "new animal" command, single line with three strings
// "new animal", name of the new animal, type of new animal
// “Created it!” on the screen after created

// Each “query” command, single line with 3 strings
// “query”, name of the animal, name of the information requested about
// Process query command by printing out the requested data

// Test with three New animal commands
// Test with three query commands
