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
	"strconv"
	"strings"
)

func swap(intI, intJ int) {

	fmt.Println("Swap function i is%v \n j is %v \n", intI, intJ)

}

func bubbleSort(intSlice []int) {

	// for _, i := range intSlice {
	// 	if intSlice[i] > intSlice[i+1] {
	// 		intI := intSlice[i]
	// 		intJ := intSlice[i+1]
	// 		swap(intI, intJ)

	// 	}
	// }

	fmt.Println("Bubble Sort function")
	fmt.Printf("You input: %v\n", intSlice)
}

func userInput(input string) []int {
	// take the string of numbers "23 57 1 576"
	// separate by the space and create a slice
	iStr := strings.Split(input, " ")
	// create an empty slice for integers
	iSlice := []int{}
	// for each index convert the string to an integer and append to iSlice
	for i := 0; i < len(iStr); i++ {
		j, _ := strconv.Atoi(iStr[i])
		iSlice = append(iSlice, j)
	}
	// at the end of the loop print the slice
	fmt.Printf("iSlice is %v\n", iSlice)
	//return iSlice
	return iSlice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a sequence of numbers, separated by a space. No more than 10.")
	//Sample input: 23 57 1 576
	scanner.Scan()
	input := scanner.Text()
	// input will have "23 57 1 576" as a string
	// send string of numbers to the userInput function to get a slice with integers for the BubbleSort function
	iSlice := userInput(input)

	//Call bubbleSort on iSlice
	bubbleSort(iSlice)

}
