package main

import "fmt"

// VARIABLE "myInt" is of INTEGER type
var myInt int

// CREATE A NEW CUSTOM TYPE
// myInteger is a custom type with underlying type as INTEGER
type myInteger int

// temp is a variable of type myInteger
var temp myInteger

func main() {
	myInt = 10
	temp = 100
	fmt.Printf("Built-in type variable :: value :: %d :: type :: %T\n", myInt, myInt)
	fmt.Printf("Custom type variable :: value :: %d :: type :: %T\n", temp, temp)

	// temp = myInt -> cannot do this, as "temp" is of type "myInteger" and "myInt" is of type INTEGER
}
