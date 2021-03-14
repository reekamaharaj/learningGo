package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program started. Enter an integer to be added and sorted.")

	intSlice := make([]int, 3)
	// count := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	for i := range intSlice {
		fmt.Println("Enter integer. Or x to exit.")

		if userInput == "X" || userInput == "x" {
			fmt.Println("You exited the program.")
			break
		}

		var intEntered int
		fmt.Scanln(&intEntered)
		intSlice[i] = intEntered
		fmt.Println("Slice: ", intSlice)
	}

}

// 	fmt.Printf("The slice you created has a length of %d a capacity of %d and the full slice is %v\n", len(intSlice), cap(intSlice), intSlice)
// }

// Code works okay. Isn't perfect at all. x will exit before the loop has started and during. But doesn't return message that program is exited. Code should also check for values and send an err when a number isn't entered...It doesn't do that currently. Still needs work.
