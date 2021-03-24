package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //bufio module (buffered input/output) create a NewScanner object as scanner. Object has od.Stdin. os: operating system. Stdin: what is input on the command line
	fmt.Printf("Type the year you were born: ")
	scanner.Scan() // will be scanned as a string
	//Scan the user input
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	//strconv will convert string to an int. at base 10 and size 64
	//save the scanned input
	// _ ignore error if there is one and continue
	fmt.Printf("At the end of 2020 you were %d years old", 2020-input)
}

//age at 2020 doesn't check for input being a number...
