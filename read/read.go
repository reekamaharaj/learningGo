package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
)

type name struct {
	fName string
	lName string
}

func main() {
	// Prompt for user to enter file name
	fmt.Println("Enter file name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	file := scanner.Text()

	// file := "names.txt"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var text []string

	for fileScanner.Scan() {
		text = append(text, fileScanner.Text())
	}

	f.Close()

	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}

// [x] prompt user for file name
// [x] read from a file
// [] series of names. first. last. separated by a space'
// [] represent text in a slice of structs
// [] define a name struct with fname and lname. string size 20
// [] New struct for each name
// [] add struct to slice
// [] print the first and last name in each struct
