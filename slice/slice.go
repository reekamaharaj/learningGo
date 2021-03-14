package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Slice started. Hit Enter to continue or type X to exit.")
	intSlice := make([]int, 3)
	// intSlice := []int{1, 2, 3}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	for i := range intSlice {
		//i is the index
		if userInput == "X" || userInput == "x" {
			fmt.Println("You exited the program.")
			break
		}

		var intEntered int
		fmt.Println("Enter integers for storage, type X to exit")
		fmt.Scanln(&intEntered)
		intSlice[i] = intEntered
	}

	fmt.Printf("The slice you created has a length of %d a capacity of %d and the full slice is %v\n", len(intSlice), cap(intSlice), intSlice)
}

// Code works okay. Isn't perfect at all. x will exit before the loop has started and during. But doesn't return message that program is exited. Code should also check for values and send an err when a number isn't entered...It doesn't do that currently. Still needs work.
