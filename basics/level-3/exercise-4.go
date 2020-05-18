/**
Create a for loop using this syntax
	for {}
Have it print out the years you have been alive
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 1995
	for {
		if birthYear > time.Now().Year() {
			break
		}
		fmt.Println(birthYear)
		birthYear++
	}
}