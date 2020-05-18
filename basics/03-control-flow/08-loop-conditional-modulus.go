package main

import "fmt"

func main() {
	fmt.Println("Printing all even numbers from 1 to 100")
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Print("\t", i)
		}
	}
	fmt.Println()
	fmt.Println("Printing all odd numbers from 1 to 100")
	for i := 1; i <= 100; i++ {
		if i % 2 != 0 {
			fmt.Print("\t", i)
		}
	}
	fmt.Println()
}
