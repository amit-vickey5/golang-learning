package main

import "fmt"

// STRING is a slice of byte
// STRING is immutable

func main() {
	str1 := "Hello there. This is first string"
	str2 := `This is second "string"
Declaring the string this way
	can take up raw values, which is not
	possible when declaring with "<string>"`
	fmt.Println("First String :: ", str1)
	fmt.Println("Second String :: ", str2)
	fmt.Printf("str1 TYPE :: %T :: str2 TYPE :: %T\n", str1, str2)

	str1 = "my string"
	fmt.Println("New value of first String :: ", str1)

	// string as a slice of byte
	bs := []byte(str1) // convert string to the byte of slice
	fmt.Println("bs value :: ", bs)
	fmt.Printf("bs TYPE :: %T\n", bs)

}
