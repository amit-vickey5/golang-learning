/**
1. Create a func with identifier foo that returns an int
2. Create a func with identifier bar that returns an int and a string
3. Call both funcs
4. Print out their result
*/

package main

import (
    "fmt"
)

func main() {
    //3...
    x1 := foo()
    x2, str := bar()

    //4...
    fmt.Println("foo() return values ::", x1)
    fmt.Println("bar() return values ::", x2, ",", str)
}

//1...
func foo() int {
    return 10
}

//2...
func bar() (int, string) {
    return 20, "Hello"
}
