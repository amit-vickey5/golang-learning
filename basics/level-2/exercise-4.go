/**
Write a program that
	a. assigns an int to a variable
	b. prints that int in decimal, binary and hex
	c. shifts the bits of that int over 1 position to the left, and assigns that to a variable
	d. prints that variable in decimal, binary and hex
 */

package main

import "fmt"

func main() {
	//a...
	num1 := 20

	//b...
	fmt.Println("Number\t:: ", num1)
	fmt.Printf("Decimal\t\t:: %d\n", num1)
	fmt.Printf("Binary\t\t:: %b\n", num1)
	fmt.Printf("Hexadecimal\t:: %#x\n\n", num1)

	//c...
	num2 := num1 << 1

	//d...
	fmt.Println("New number\t:: ", num2)
	fmt.Printf("Decimal\t\t:: %d\n", num2)
	fmt.Printf("Binary\t\t:: %b\n", num2)
	fmt.Printf("Hexadecimal\t:: %#x\n", num2)
}