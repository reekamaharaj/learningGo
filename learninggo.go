package main //package declaration

import (
	"bufio"
	"strconv"

	//buffered input/output
	//creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

	"fmt"
	// import format package, implements formatting for input and output
	// format 'verbs' derived from C, but simpler. %q: single-quoted character literal. %d base 10, %T type value %v value in a default format. Print with Printf

	"math"
	//import math package
	"os"
	//Package os provides a platform-independent interface to operating system functionality.
	"strings"
	//import strings package
	//"strconv"
	//Package strconv implements conversions to and from string representations of basic data types.
)

func main() {
	trunc()
} //function with no parameters, the main function will get called when the program is executed.

func helloWorld() {
	fmt.Println("hello world")
}

// trunc is a function which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.

func trunc() {
	var floatNum float64
	fmt.Println("Enter a floating point")

	fmt.Scanln(&floatNum)
	fmt.Println(math.Trunc(floatNum))
}

// findIan is a function that will prompt a user to enter a string and will check the input to see if it starts with i, ends with n and contains a. Case does not matter, spaces do not matter. Found will be returned on success. Not Fount will be returned otherwise.

// convert the input to lower case to account for capital letters.
// spaces are a problem for the strings package, needed bufio package to handle spaces

// Test cases:
// Found: “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N
// Not Found: “ihhhhhn”, “ina”, “xian”

func findIan() {
	fmt.Println("Enter a string") //instruction to user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	lowerString := strings.ToLower(userInput)
	startsWithI := strings.HasPrefix(lowerString, "i")
	endsWithN := strings.HasSuffix(lowerString, "n")
	hasA := strings.Contains(lowerString, "a")

	if startsWithI == true && endsWithN == true && hasA == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found")
	}
}

// slice is a function which prompts the user to enter integers and stores the integers in a sorted slice. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program will only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

func slice() {
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

}

func ageAt2020() {
	scanner := bufio.NewScanner(os.Stdin) //bufio module (buffered input/output) create a NewScanner object as scanner. Object has od.Stdin. os: operating system. Stdin: what is input on the command line
	fmt.Printf("Type the year you were born: ")
	scanner.Scan() // will be scanned as a string
	//Scan the user input
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	//strconv will convert string to an int. at base 10 and size 64
	//save the scanned input
	// _ ignore error if there is one and continue
	fmt.Printf("At the end of 2020 you were %d years old", 2020-input)
}

//age at 2020 doesn't check for input being a number...
