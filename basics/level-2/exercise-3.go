/**
Create TYPED and UNTYPED constants. Print the values of the constants.
 */

package main

import "fmt"

const (
	typedConstant int = 10
	untypedConstant = 20
)

func main() {
	fmt.Println("Typed Constant :: ", typedConstant)
	fmt.Println("Untyped Constant :: ", untypedConstant)
}
