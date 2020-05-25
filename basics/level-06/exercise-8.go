/**
1. create a func which returns a func
2. assign the returned func to a variable
3. call the retuned func
*/

package main

import (
    "fmt"
)

func main() {
    //2...
    returnedFunc := foo()

    //3...
    returnedFunc()
}

//1...
func foo() func() {
    return func() {
        fmt.Println("returned func")
    }
}
