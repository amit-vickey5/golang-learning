package main

import "fmt"

// UNTYPED CONSTANTS...here compiler decides the type of the constant
const myIntConstant1 = 10
const (
	myFloatConstant1 = 24.42
	myStringConstant1 = "Amit"
)

// TYPED CONSTANTS
const (
	myIntConstant2 int = 20
	myFloatConstant2 float64 = 44.42
	myStringConstant2 string = "Vickey"
)

func main() {
	fmt.Println("UNTYPED CONSTANTS :: ")
	fmt.Printf("Constant 'myIntConstant1' :: value %v :: type :: %T\n", myIntConstant1, myIntConstant1)
	fmt.Printf("Constant 'myFloatConstant1' :: value %v :: type :: %T\n", myFloatConstant1, myFloatConstant1)
	fmt.Printf("Constant 'myStringConstant1' :: value %v :: type :: %T\n", myStringConstant1, myStringConstant1)
	fmt.Println("TYPED CONSTANTS :: ")
	fmt.Printf("Constant 'myIntConstant2' :: value %v :: type :: %T\n", myIntConstant2, myIntConstant2)
	fmt.Printf("Constant 'myFloatConstant2' :: value %v :: type :: %T\n", myFloatConstant2, myFloatConstant2)
	fmt.Printf("Constant 'myStringConstant2' :: value %v :: type :: %T\n", myStringConstant2, myStringConstant2)
}
