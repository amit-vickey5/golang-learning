/**
Create a program that uses switch statement with no switch expression specified
 */

package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Hi...this is TRUE block")
	case false:
		fmt.Println("Hi...this is FALSE block")
	default:
		fmt.Println("Hi...this is DEFAULT block")
	}
}