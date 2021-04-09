// time t, acceleration a, initial velocity vo, initial displacement so
// s = ½ a t2 + vot + so

// Prompt user for a, vo, so
// Prompt user to enter time
// Program will compute the displacement after the time

//acceleration, initial velocity, initial displacement, time
//Function called GenDisplaceFn(a, vo, so float64) func (t float64) float64 {}

// GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so
// GenDisplaceFn() should return a function which computes displacement as a function of time
// function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t

// a = 10, vo = 2, so = 1
// fn := GenDisplaceFn(10, 2, 1)
// print the displacement after 3 seconds.
// fmt.Println(fn(3))
// the displacement after 5 seconds.
// fmt.Println(fn(5))

//Test. Run twice with two sets of values

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(input string) []float64 {
	iStr := strings.Split(input, "")
	iSlice := []float64{}
	for i := 0; i < len(iStr); i++ {
		j, _ := strconv.Atoi(iStr[i])
		iSlice = append(iSlice, j)
	}
	fmt.Printf("iSlice is %v\n", iSlice)
	return iSlice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an acceleration, initial velocity and initial displacement. Separate with a space and the values can be a float.")
	scanner.Scan()
	input := scanner.Text()
	aVoSoVals := userInput(input)
}
