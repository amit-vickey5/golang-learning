package main

import "fmt"

func main() {
	// if statement with bool
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}

	// if statement with !(not) operator
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}

	// if statement with expression
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}

	// if statement with initialization statement
	// syntax: if initialization; condition {}
	// here the scope of "x" is only in the if-block
	if x := 42; x < 50 {
		fmt.Println("007")
	}
	// fmt.Println(x) -> cannot do, as the scope of "x" is only in the if-block

	if y := 42; y > 50 {
		fmt.Println("008")
	}
}
