package main

import "fmt"

// If you use "int" as the variable type, GoLang will automatically decide whether to
// use int32 or int64 based on the architecture of the system. Same goes with uint

// "byte" is an alias for "int8"

func main() {
	x := 24
	y := 42.2345
	fmt.Println("Value of 'x' :: ", x)
	fmt.Printf("Type of 'x' :: %T\n", x)
	fmt.Println("Value of 'y' :: ", y)
	fmt.Printf("Type of 'y' :: %T\n", y)
}
