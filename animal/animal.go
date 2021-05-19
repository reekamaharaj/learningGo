package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func main() {
	//Cow struct
	cow := Animal{
		name:       "Cow",
		food:       "Grass",
		locomotion: "Walk",
		noise:      "Moo",
	}

	//Bird struct
	bird := Animal{
		name:       "Bird",
		food:       "Worms",
		locomotion: "Fly",
		noise:      "Peep",
	}

	//Snake struct
	snake := Animal{
		name:       "Snake",
		food:       "Mice",
		locomotion: "Slither",
		noise:      "Hiss",
	}
	for {
		fmt.Println("Choose an animal: Cow, Bird, or Snake.")
		fmt.Println("Choose a topic: Food, Locomotion, or Noise.")
		fmt.Println("Enter both. For a response about the animal.\n Example if you entered: Dog Noise \n You would get a response with 'Bark'.")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		lowerString := strings.ToLower(userInput)
		s := strings.Split(lowerString, " ")
		animal, action := s[0], s[1]

		switch animal {
		case "cow":
			if action == "food" {
				fmt.Println("\n Response:", cow.food, "\n")
			}

			if action == "locomotion" {
				fmt.Println("\n Response:", cow.locomotion, "\n")
			}

			if action == "noise" {
				fmt.Println("\n Response:", cow.noise, "\n")
			} else {
				fmt.Println("\n Something went wrong. Please try again. \n")
			}

		case "bird":
			if action == "food" {
				fmt.Println("\n Response:", bird.food, "\n")
			}

			if action == "locomotion" {
				fmt.Println("\n Response:", bird.locomotion, "\n")
			}

			if action == "noise" {
				fmt.Println("\n Response:", bird.noise, "\n")
			} else {
				fmt.Println("\n Something went wrong. Please try again. \n")
			}
		case "snake":
			if action == "food" {
				fmt.Println("\n Response:", snake.food, "\n")
			}

			if action == "locomotion" {
				fmt.Println("\n Response:", snake.locomotion, "\n")
			}

			if action == "noise" {
				fmt.Println("\n Response:", snake.noise, "\n")
			} else {
				fmt.Println("\n Something went wrong. Please try again. \n")
			}
		default:
			fmt.Println("\n Something went wrong. Please try again. \n")
		}
	}

}

// Get info about animals
//Defined animals
//Cow, bird, Snake
//Defined actions
//Eat, Move, Speak
//Grass, Worms, Mice
//Walk, Fly, Slither
//Moo, Peep, Hiss

//User can for information about animals
//What food it eats
//How it moves
//What sound it makes

//">" promt to user to type request
//One request at a time with two strings
//first string animal. second stringg action
//returns the answer
//loops forever

// You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.

// Submit your Go program source code.

// This assignment is worth a total of 10 points.

// Test the program by running it and testing it by issuing three requests. Each request should involve a different animal (cow, bird, snake) and a different property of the animal (eat, move, speak). Give 2 points for each request for which the program provides the correct response.

// Examine the code. If the code contains a type called Animal, which is a struct containing three fields, all of which are strings, then give another 2 points. If the program contains three methods called Eat(), Move(), and Speak(), and all of their receiver types are Animal, give another 2 points.
