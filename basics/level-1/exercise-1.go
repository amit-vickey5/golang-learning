/**
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
IDENTIFIERS "x" and "y" and "z"
	a. 42
	b. "James Bond"
	c. true
2. Now print the values stored in those variables using
	a. a single print statement
	b. multiple print statements
 */

package main

import "fmt"

func main() {
	//1...
	x := 42
	y := "James Bond"
	z := true

	//2(a)...
	fmt.Println(x, y, z)

	//2(b)...
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
