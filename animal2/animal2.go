package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserAnimal struct {
	name       string
	animalType string
}

func (u UserAnimal) NewAnimal() {
	fmt.Println("\n Response:", u.name, "\n")
}

type Animal interface {
	Eat()
	Move()
	Speak()
	Name() string
}

// Cow struct and methods
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("\n Response: Grass \n")
}

func (c Cow) Move() {
	fmt.Println("\n Response: Walk \n")
}

func (c Cow) Speak() {
	fmt.Println("\n Response: Moo \n")
}

// Bird struct and methods
type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("\n Response: Worms \n")
}

func (b Bird) Move() {
	fmt.Println("\n Response: Fly \n")
}

func (b Bird) Speak() {
	fmt.Println("\n Response: Hiss \n")
}

// Snake struct and methods
type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("\n Response: Mice \n")
}

func (s Snake) Move() {
	fmt.Println("\n Response: Slither \n")
}

func (s Snake) Speak() {
	fmt.Println("\n Response: Peep \n")
}

func main() {
	fmt.Println("Would you like to make a query about an animal or add a new animal? Enter 'query' or 'new' to get started.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	command := strings.ToLower(userInput)

	switch command {
	case "query":
		fmt.Println("What would you like to know?")
		fmt.Println("Choose an animal: Cow, Bird, or Snake.")
		fmt.Println("Choose a topic: Food, Locomotion, or Noise.")
		fmt.Println("Enter both. For a response about the animal.\n Example if you entered: Dog Noise \n You would get a response with 'Bark'.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		lowerString := strings.ToLower(userInput)
		s := strings.Split(lowerString, " ")
		animal, action := s[0], s[1]
	case "new":
		fmt.Println("Enter a name and an animal type (cow, bird or snake)")

	default:
		fmt.Println("\n Something went wrong. Please try again.")
	}
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
