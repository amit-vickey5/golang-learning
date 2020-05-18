package main

import "fmt"

func main() {
	// For statement with single sondition
	fmt.Println("Single Condition :: ")
	x := 1
	for x < 10 {
		fmt.Println("\t", x)
		x++
	}
	fmt.Println("Done")

	// Eternal(infinite loop)
	fmt.Println("Infinite Loop :: ")
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println("\t", y)
		y++
	}
	fmt.Println("Done")

}
