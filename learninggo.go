package main //package declaration

import (
	"bufio"
	"fmt"  // import format package, implements formatting for input and output
	"math" //import math package
	"os"
	"strings" //import strings package
)

func main() {
	findIan()
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
