/**
Exercise 2 ::
	1. Use var to DECLARE three variables. The variables should have package level scope.
	Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
	variables and make sure the variables are of the following TYPE(meaning they can
	store VALUES of that TYPE)
		a. identifier "x" type int
		b. identifier "y" type string
		c. identifier "z" type bool
	2. in func main
		a. print out the values for each variable
		b. The compiler assigned values to these variables. What are these values called?

Exercise 3 ::
	Using the code from the previous exercise,
	1. At the package level scope, assign the following values to the three variables
		a. for x assign 42
		b. for y assign "James Bond"
		c. for z assign true
	2. in func main()
		a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
		returned value of TYPE string using the short declaration operator to a
		VARIABLE with the IDENTIFIER "s"
		b. print out the value stored by variable "s"
 */

package main

import "fmt"

//2-1...3-1...
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//2-2(a)...
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//2-2(b)...ZERO(or default) VALUES

	//3-2(a)...
	s := fmt.Sprintf("%d\t%s\t%t", x, y, z)

	//3-2(b)...
	fmt.Println(s)
}