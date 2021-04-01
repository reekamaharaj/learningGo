package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
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

	// scan file by line and add to slice
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var names []*Name

	for fileScanner.Scan() {
		fullName := fileScanner.Text()
		s := strings.Split(fullName, " ")
		fName, lName := s[0], s[1]
		names = append(names, &Name{fName: fName, lName: lName})
	}

	for _, value := range names {
		fmt.Printf("%#v\n", value)
	}

	f.Close()

}

// [x] prompt user for file name
// [x] read from a file
// [x] series of names. first. last. separated by a space'
// [x] represent text in a slice of structs
// [x] define a name struct with fname and lname.
// [] string size 20 for each name
// [x] New struct for each name
// [x] add struct to slice
// [x] print the first and last name in each struct
