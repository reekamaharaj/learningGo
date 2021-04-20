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
	"math"
	"os"
	"strconv"
)

func input(input string) float64 {
	i, _ := strconv.ParseFloat(input, 64)
	return i
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((0.5 * a * (math.Pow(t, 2))) + (vo * t) + so)
	}
}

func main() {

	aScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an acceleration")
	aScanner.Scan()
	aInput := aScanner.Text()
	a := input(aInput)

	vScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter initial velocity")
	vScanner.Scan()
	vInput := vScanner.Text()
	v := input(vInput)

	sScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter initial displacement")
	sScanner.Scan()
	sInput := sScanner.Text()
	s := input(sInput)

	tScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a time in seconds")
	tScanner.Scan()
	tInput := tScanner.Text()
	t := input(tInput)

	fmt.Printf("Acceleration = %v, Initial velocity = %v, Initial displacement = %v, Time = %v\n", a, v, s, t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println("The displacement will be calculated for 3 seconds")
	fmt.Println(fn(t))
}
