package main

import "fmt"

func main() {
	// if-else
	x := 42
	fmt.Println("if-else block :: ")
	if x == 40 {
		fmt.Println("\tOur value was 40")
	} else {
		fmt.Println("\tOur value was not 40")
	}

	// chain of if-else
	fmt.Println("Chain of if-else :: ")
	if x == 40 {
		fmt.Println("\tOur value was 40")
	} else if x == 41 {
		fmt.Println("\tOur value was 41")
	} else if x == 42 {
		fmt.Println("\tOur value was 42")
	} else if x == 43 {
		fmt.Println("\tOur value was 43")
	} else {
		fmt.Println("\tOur value was not 40, 41, 42, or 43")
	}

}
