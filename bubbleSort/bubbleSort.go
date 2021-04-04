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

func swap(intSlice []int) {
	//first number and second number, compare
	// if first > second then swap first and second
	// if first < second then do nothing and contiue the loop
	//first > second so. first = second and second = first
	var i int = 0
	var j int = 1

	for j < len(intSlice) {
		var iInt int = intSlice[i]
		var jInt int = intSlice[j]

		if iInt > jInt {
			intSlice[i] = jInt
			intSlice[j] = iInt
		}
		i++
		j++
	}
}

func bubbleSort(intSlice []int) []int {
	// going through the whole intSlice slice
	for i := 0; i < len(intSlice); i++ {
		swap(intSlice)
	}
	return intSlice
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
	//Even numbers: 22 10 38 2 100 62
	//With a negative number: -2 547 2 -4 38
	scanner.Scan()
	input := scanner.Text()
	// input will have "23 57 1 576" as a string
	// send string of numbers to the userInput function to get a slice with integers for the BubbleSort function
	iSlice := userInput(input)

	//Call bubbleSort on iSlice
	fmt.Printf("Sorted slice %v", bubbleSort(iSlice))

}
