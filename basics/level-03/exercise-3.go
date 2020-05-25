/**
Create a for loop using this syntax
	for condition {}
Have it print out the years you have been alive
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 1995
	for birthYear <= time.Now().Year() {
		fmt.Println(birthYear)
		birthYear++
	}
}
