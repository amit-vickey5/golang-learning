/**
Write a program that prints a number in decimal, binary and hex
 */

package main

import "fmt"

func main() {
	number := 42
	fmt.Println("Number\t:: ", number)
	fmt.Printf("Decimal\t\t:: %d\n", number)
	fmt.Printf("Binary\t\t:: %b\n", number)
	fmt.Printf("Hexadecimal\t:: %#x\n", number)
}
