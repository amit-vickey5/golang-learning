/**
Using iota, create 4 constants for the last 4 years. Print the constant values
 */

package main

import (
	"fmt"
)

const (
	currYear = 2020 + iota
	year1 = currYear - iota
	year2 = currYear - iota
	year3 = currYear - iota
)

func main() {
	fmt.Println("Current Year :: ", currYear)
	fmt.Println("Year - 1 :: ", year1)
	fmt.Println("Year - 2 :: ", year2)
	fmt.Println("Year - 3 :: ", year3)
}