// prompt the user to type in a sequence of up to 10 integers.
// print the integers out on one line, in sorted order, from least to
// greatest.
// function called BubbleSort()
// takes a slice of integers as an argument and returns nothing
//BubbleSort() function should modify the slice

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements
// write a Swap() function which performs this operation
// function should take two arguments
// a slice of integers and an index value i which
// indicates a position in the slice
// return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

// Test: sequence of integers should be all positive numbers
// Test: at least one negative number

package main

import (
	"bufio"
	"fmt"
	"os"
)

func swap() {
	fmt.Println("Swap fucntion")
}

func bubbleSort() {
	fmt.Println("Bubble Sort function")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a sequence of numbers. No more than 10.")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You input: %v\n", input)

	swap()

	bubbleSort()

}
