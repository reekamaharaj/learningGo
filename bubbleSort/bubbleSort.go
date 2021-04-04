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

	for _, i := range intSlice {
		if intSlice[i] > intSlice[i+1] {
			intI := intSlice[i]
			intJ := intSlice[i+1]
			swap(intI, intJ)

		}
	}

	fmt.Println("Bubble Sort function")
	fmt.Printf("You input: %v\n", intSlice)
}

func userInput(input string) []int {
	iStr := strings.Split(input, " ")

	iSlice := []int{}

	for _, i := range iStr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iSlice = append(iSlice, j)
	}
	fmt.Printf("iSlice is %v\n", iSlice)
	return iSlice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a sequence of numbers, separated by a space. No more than 10.")

	scanner.Scan()
	input := scanner.Text()
	iSlice := userInput(input)

	bubbleSort(iSlice)

}
