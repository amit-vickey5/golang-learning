package main

import "fmt"

func main() {
	//DECLARE a BOOLEAN variable
	var x bool
	fmt.Println("Zero value of BOOLEAN :: ", x)
	x = true
	fmt.Println("Modified value of 'x' :: ", x)

	a := 7
	b := 42
	fmt.Println("Some boolean operators ::")
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a >= b", a >= b)
}
