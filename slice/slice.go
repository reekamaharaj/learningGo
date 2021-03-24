package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	//program starts
	fmt.Println("The slice sort program has started.")
	//create an empty integer slice of length 3
	intSlice := make([]int, 3)
	//starts a counter
	count := 0

	//loop starts
	for {
		//Tells program to scan input using bufio module, os: operating system, Stdin: standard input, what is input to command line. Creates object scanner.
		scanner := bufio.NewScanner(os.Stdin)

		//Prompt user to enter integer
		fmt.Println("Enter a number to add to the slice. Hit X to exit program")

		//scan what the user input as a string
		scanner.Scan()

		//Save the text scanned to the variable userInput, will be saved as a string
		userInput := scanner.Text()

		//check if x or X is entered and exit if yes
		if userInput == "X" || userInput == "x" {
			fmt.Println("You exited the program.")
			break
		} else {
			//if not exited

			//create an integer variable
			var toInt int

			//convert string to int, or return an error
			//string - userInput to int - toInt
			toInt, err := strconv.Atoi(userInput)

			// what to do if there is an error
			if err != nil {
				fmt.Println("Does not compute. Enter an integer or request to exit.")
				continue
			}
			//if the counter is less than the length of the slice (slice starts at a length of 3 and count starts at 0)
			if count < len(intSlice) {
				//then at index [count] of the slice, add the integer
				intSlice[count] = toInt

			} else {
				//if count is now equal to the length of the slice
				intSlice = append(intSlice, toInt)
			}
			count++
			sorted := intSlice[0:count]
			sort.Ints(sorted)
			fmt.Printf("The sorted array is %v \n", sorted)

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

// slice is a function which prompts the user to enter integers and stores the integers in a sorted slice. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program will only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
