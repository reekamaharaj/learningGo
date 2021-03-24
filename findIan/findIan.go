package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// findIan is a function that will prompt a user to enter a string and will check the input to see if it starts with i, ends with n and contains a. Case does not matter, spaces do not matter. Found will be returned on success. Not Fount will be returned otherwise.

// convert the input to lower case to account for capital letters.
// spaces are a problem for the strings package, needed bufio package to handle spaces

// Test cases:
// Found: “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N
// Not Found: “ihhhhhn”, “ina”, “xian”

func main() {
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
