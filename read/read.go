package main

type name struct {
	fName string
	lName string
}

func newName(fName string) *name {
	n := name{fName: fName}
	n.lName = lName
}

func main() {

}

// prompt user for file name
// read from a file
// series of names. first. last. separated by a space'
// represent text in a slice of structs
// define a name struct with fname and lname. string size 20
// New struct for each name
// add struct to slice
// print the first and last name in each struct
