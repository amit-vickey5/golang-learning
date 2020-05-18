/**
Create a program that uses "else-if" and "else"
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 1995
	curentYear := time.Now().Year()
	if curentYear - birthYear < 18 {
		fmt.Println("You are not eligible for anything")
	} else if curentYear - birthYear < 21 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are eligible to vote and marry")
	}
}