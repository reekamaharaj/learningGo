// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//Prompt user to enter name and address
	fmt.Println("Enter your first name")

	//Tells program to scan input using bufio module, os: operating system, Stdin: standard input, what is input to command line. Creates object scanner.
	scanner := bufio.NewScanner(os.Stdin)

	//scan what the user input as a string
	scanner.Scan()

	//Save the text scanned to the variable userInput, will be saved as a string
	name := scanner.Text()
	fmt.Println("Enter your address")
	addScanner := bufio.NewScanner(os.Stdin)
	addScanner.Scan()
	address := addScanner.Text()

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	mjson, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error", err)
	}
	mjsonstr := string(mjson)
	fmt.Println(mjsonstr)
}
