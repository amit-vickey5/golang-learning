package main

import "fmt"

// GoLang is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language like javascript

func main() {

	// DECLARE the VARIABLE with he IDENTIFIER "a"
	// "b" is of TYPE INT
	var a = 10

	// DECLARE the VARIABLE with he IDENTIFIER "b"
	// "b" is of TYPE STRING
	var b = "Hello there"
	fmt.Printf("Type of 'a' :: %T :: value of 'a' :: %d\n", a, a)
	fmt.Printf("Type of 'b' :: %T :: value of 'b' :: %s\n", b, b)

	// b = 42 -> cannot do as "b" is of type STRING

}
