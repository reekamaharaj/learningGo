package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("\nGrass\n")
}

func (c Cow) Move() {
	fmt.Println("\nWalk\n")
}

func (c Cow) Speak() {
	fmt.Println("\nMoo\n")
}

func (b Bird) Eat() {
	fmt.Println("\nWorms\n")
}

func (b Bird) Move() {
	fmt.Println("\nFly\n")
}

func (b Bird) Speak() {
	fmt.Println("\nPeep\n")
}

func (s Snake) Eat() {
	fmt.Println("\nMice\n")
}

func (s Snake) Move() {
	fmt.Println("\nSlither\n")
}

func (s Snake) Speak() {
	fmt.Println("\nHiss\n")
}

func main() {
	m := make(map[string]Animal)
	cow := Cow{}
	bird := Bird{}
	snake := Snake{}

	for {
		fmt.Println("To name an animal enter 'newanimal <enter animal's name> <enter type of animal>")
		fmt.Println("To query an animal enter 'query <enter animal's name> <enter action you want to query>\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		lowerString := strings.ToLower(userInput)
		s := strings.Split(lowerString, " ")
		command, name, definition := s[0], s[1], s[2]

		switch command {
		case "newanimal":
			switch definition {
			case "cow":
				m[name] = cow
				fmt.Println("\nCreated!\n")
			case "bird":
				m[name] = bird
				fmt.Println("\nCreated!\n")
			case "snake":
				m[name] = snake
				fmt.Println("\nCreated!\n")
			default:
				fmt.Println("\nSomething went wrong. Please try again.\n")
			}
		case "query":
			animal, exists := m[name]
			if !exists {
				fmt.Println("\nSomething went wrong. Please try again.\n")
			} else {
				switch definition {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("\nSomething went wrong. Please try again.\n")
				}
			}
		default:
			fmt.Println("\nSomething went wrong. Please try again.\n")
		}
	}
}
