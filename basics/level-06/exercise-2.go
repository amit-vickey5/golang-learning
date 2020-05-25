/**
1. Create a func with identifier foo that
    - takes a variadic parameter of type int
    - pass in a value of type []int into your function
    - returns the sum of all the values of type int passed in
2. Create a func  with the identifier bar that
    - takes in a parameter of type []int
    - returns the sum of all the values of type int passed
*/

package main

import (
    "fmt"
)

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fooSum := foo(arr...)
    barSum := bar(arr)
    fmt.Println("Elements ::", arr)
    fmt.Println("Sum by foo() ::", fooSum)
    fmt.Println("Sum by bar() ::", barSum)
}

//1...
func foo(arr ...int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

//2...
func bar(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
