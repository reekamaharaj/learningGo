package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	month := time.Month(1 + rand.Intn(12))
	year := (2011 + rand.Intn(10))

	fmt.Printf("month = %v\n", month)
	fmt.Printf("year = %v\n", year)

}
