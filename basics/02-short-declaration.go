package main

import "fmt"

func main() {
	x := 42 //Declare and Initialize the variable x
	fmt.Println("Initial value of x :: ", x)
	x = 100 //Change the value of the already declared variable x
	fmt.Println("New value of x :: ", x)
	y := 10 + x //Initial value can also be given as an expression
	fmt.Println("Value of y :: ", y)
}
