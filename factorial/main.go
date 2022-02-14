package main

import "fmt"

var output int

// func factorial(number int) int {
// 	if number <= 1 {
// 		return 1
// 	}
// 	return number * factorial(number-1)
// }

// func main() {
// 	fmt.Println(factorial(4))
// }

func fact(number int) {
	var output int = 1
	if number <= 0 {
		fmt.Println("0 or lower")
	} else {
		for i := 1; i <= number; i++ {
			output = i * output
		}
		fmt.Println(output)
	}
}
func main() {
	fact(3)
	fact(5)
}
