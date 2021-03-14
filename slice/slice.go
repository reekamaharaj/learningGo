package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//program starts
	fmt.Println("The slice sort program has started.")
	//create an empty integer slice of length 3
	intSlice := make([]int, 3)
	count := 0

	//loop start
	for {
		scanner := bufio.NewScanner(os.Stdin)
		//Tells program to scan input using bufio module, os: operating system, Stdin: standard input, what is input to command line. Creates object scanner.
		fmt.Println("Enter a number to add to the slice. Hit X to exit program")
		//Prompt user to enter integer
		scanner.Scan()
		//scan what the user input as a string
		userInput := scanner.Text()
		//Save the text scanned to the variable userInput

		//check if x or X is entered and exit if yes
		if userInput == "X" || userInput == "x" {
			fmt.Println("You exited the program.")
			break
		} else {
			var toInt int
			//convert string to int, or return an error
			toInt, err := strconv.Atoi(userInput)
			// what to do if there is an error
			if err != nil {
				fmt.Println("Does not compute. Enter an integer or request to exit.")
				continue
			}
			if count < len(intSlice) {
				intSlice[count] = toInt
				fmt.Printf("The length of the slice is %d and it contains %v \n", len(intSlice), intSlice)
				count++
			} else {
				intSlice = append(intSlice, toInt)
				fmt.Printf("The length of the slice is %d and it contains %v \n", len(intSlice), intSlice)
				count++
			}

		}

	}

}

//if an integer is added then add it to the slice
//sort the slice
//print the sorted slice
//ask for another integer
//end loop. or continue with the loop.

//slice has to grow in size if the user continues to add integers.

//Does it compile
//Does it print a sorted slice after every entry
//Can more than three integers be added and sorted?

//Should check that the entered value is a number, and if it isn't ask user for new input that is a number
