/**
Use the following operators, write expressions and assign their values to variables:
	a. ==
	b. <=
	c. >=
	d. !=
	e. <
	f. >
 */

package main

import "fmt"

func main() {
	number1 := 10
	number2 := 20
	a := number1 == number2
	b := number1 <= number2
	c := number1 >= number2
	d := number1 != number2
	e := number1 < number2
	f := number1 > number2
	fmt.Println("a...", a)
	fmt.Println("b...", b)
	fmt.Println("c...", c)
	fmt.Println("d...", d)
	fmt.Println("e...", e)
	fmt.Println("f...", f)
}
