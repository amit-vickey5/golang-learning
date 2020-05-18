/**
Create a program that shows an if statement in action
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 1995
	curentYear := time.Now().Year()
	if curentYear - birthYear >= 18 {
		fmt.Println("Congrats, you are eligible to have sex.")
	} else {
		fmt.Println("Grow a dick first kid.!!!!")
	}
}