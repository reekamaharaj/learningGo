package main //package declaration

import (
	"bufio"
	"fmt"  // import format package, implements formatting for input and output
	"math" //import math package
	"os"
	"strings" //import strings package
)

func main() {
	slice()
} //function with no parameters, the main function will get called when the program is executed.

func helloWorld() {
	fmt.Println("hello world")
}

// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

func trunc() {
	var floatNum float64
	fmt.Println("Enter a floating point")

	fmt.Scanln(&floatNum)
	fmt.Println(math.Trunc(floatNum))
}

// convert the input to lower case
// first letter must be i, last letter must be n, and an a needs to be in the middle somewhere. - found case. otherwise not found
// Test for: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

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

func slice() {

	// intSlice := make([]int, 3)
	intSlice := []int{1, 2, 3}
	for i := range intSlice {
		//i is the index
		var intEntered int
		fmt.Println("Enter integers for storage")
		fmt.Scanln(&intEntered)
		intSlice[i] = intEntered
	}
	fmt.Printf("The slice you created has a length of %d a capacity of %d and the full slice is %v\n", len(intSlice), cap(intSlice), intSlice)
}

// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
