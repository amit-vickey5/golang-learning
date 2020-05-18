package main

import "fmt"

func main() {

	fmt.Println("switch statement")
	switch {
	case (2 == 2):
		fmt.Println("\tThis should print: 001")
	case false:
		fmt.Println("\tThis should not print: 001")
	}

	fmt.Println("switch with fallthrough")
	// if the true case has "fallthrough", the case immediately after that true case works(even if it's a false case)
	switch {
	case (2 == 2):
		fmt.Println("\this should print: 002")
		fallthrough // will execute the next case even if it is false
	case false:
		fmt.Println("\tThis should not print: 002")
		fallthrough // will execute the next case i.e. default case as well
	default:
		fmt.Println("\tDefault case")
	}

	fmt.Println("switch on a value")
	myText := "X"
	switch myText {
	case "A":
		fmt.Println("\tThe text is 'A'")
	case "B":
		fmt.Println("\tThe text is 'B'")
	case "C":
		fmt.Println("\tThe text is 'C'")
	case "X":
		fmt.Println("\tThe text is 'X'")
	case "Y":
		fmt.Println("\tThe text is 'Y'")
	case "Z":
		fmt.Println("\tThe text is 'Z'")
	default:
		fmt.Println("\tDefault Case")
	}

	fmt.Println("switch with cases separated as comma separated list")
	switch myText {
	case "A", "B", "C":
		fmt.Println("\tThe text is 'A' or 'B' or 'C'")
	case "X", "Y", "Z":
		fmt.Println("\tThe text is 'X' or 'Y' or 'Z'")
	default:
		fmt.Println("\tDefault Case")
	}
}
