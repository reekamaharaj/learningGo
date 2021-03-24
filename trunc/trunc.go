package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNum float64
	fmt.Println("Enter a floating point")

	fmt.Scanln(&floatNum)
	fmt.Println(math.Trunc(floatNum))
}

// trunc is a function which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
