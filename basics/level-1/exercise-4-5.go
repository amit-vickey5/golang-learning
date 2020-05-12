/**
Exercise 4 ::
	1. Create your own type. Have the underlying type to be an int.
	2. Create a VARIABLE of your new TYPE with the IDENTIFIER "a" using the "VAR" keyword
	3. in func main
		a. print out the value of the variable "a"
		b. print out the type of the variable "a"
		c. assign 42 to the variable "a" using the "=" OPERATOR
		d. print out the value of the variable "a"

Exercise 5 ::
	1. at the package level scope, using the "var" keyword, create a VARIABLE with the
	IDENTIFIER "b". The variable should be of the UNDERLYING TYPE of your custom TYPE "a"
	2. in func main
		a. use CONVERSION to convert the TYPE of the VALUE stored in "a" to the
		UNDERLYING TYPE
		b. then use the short declaration operator to ASSIGN that value to "b"
		c. print out the value stored in "b"
		d. print out the type of "b"
 */

package main

import "fmt"

//4-1...
type myCustomIntegerType int

//4-2...
var a myCustomIntegerType

//5-1
var b int

func main() {
	//4-3(a)...
	fmt.Println("Value of 'a' :: ", a)

	//4-3(b)...
	fmt.Printf("Type of 'a' :: %T\n", a)

	//4-3(c)...
	a = 42

	//4-3(d)...
	fmt.Println("Value of 'a' :: ", a)

	//5-2(a)...5-2(b)...
	b = int(a)

	//5-2(c)...
	fmt.Println("Value of 'b' :: ", b)

	//5-2(d)...
	fmt.Printf("Type of 'b' :: %T\n", b)
}