package main

import (
	// "bufio"
	"fmt"
	"os"
)

type name struct {
	fName string
	lName string
}

func main() {
	// fmt.Println("Enter file name")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// file := scanner.Text()

	file := "names.txt"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	name := make([]byte, 20)
	count, err := f.Read(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, name[:count])
	f.Close()
}

// prompt user for file name
// read from a file
// series of names. first. last. separated by a space'
// represent text in a slice of structs
// define a name struct with fname and lname. string size 20
// New struct for each name
// add struct to slice
// print the first and last name in each struct
