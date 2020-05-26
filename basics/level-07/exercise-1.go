/**
1. create a value and assign it to a variable
2. print the address of that value
*/

package main

import (
    "fmt"
)

func main() {
    x := 42
    fmt.Println("Value of 'x' ::", x)
    fmt.Println("Address of 'x' ::", &x)
}
