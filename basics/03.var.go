package main

import "fmt"

// DECLARE a VARIABLE with IDENTIFIER "a"
// SCOPE is throughout the package
var a = 20

// DECLARE a VARIABLE with IDENTIFIER "b"
// TYPE of "b" is INT
// ASSIGNS the ZERO VALUE of type INT to "b"
// ZERO VALUES :: 0 for INTEGERS, 0.0 for FLOATS, "" for STRINGS
// ZERO VALUES :: nil for POINTERS, FUNCTIONS, INTERFACES, SLICES, CHANNELS, and MAPS
// SCOPE is throughout the package
var b int

// x := 10 -> SHORT DECLARATION cannot be used to DECLARE a VARIABLE at package level

func main() {
	c := 10 //Short Declaration :: DECLARES and INITIALIZES the value of x
	var d = 100
	fmt.Println("Value of a :: ", a)
	fmt.Println("Value of b :: ", b)
	fmt.Println("Value of c :: ", c)
	fmt.Println("Value of d :: ", d)
	foo()
}

func foo() {
	// fmt.Println("foo :: value of c :: ", c) -> "c" cannot be used here. scope of "c" is just within the main()
	// fmt.Println("foo :: value of d :: ", d) -> "d" also cannot be used here. scope of "d" is just within the main()
	// VARIABLES with IDENTIFIERS "a" and "b" can be used here, as they are defined at package level
	fmt.Println("foo :: value of a :: ", a)
	fmt.Println("foo :: value of b :: ", b)
}
