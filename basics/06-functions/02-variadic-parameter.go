package main

import (
    "fmt"
)

// Variadic means zero or more i.e. variadic function can be called with zero or more arguments
// Variadic parameter have to be the last parameter in the function
// If no argument is passed to the variadic function, the value passed will be "nil"

func main() {
    fmt.Println("Calling Variadic Function with some arguments :: ")
    sum := foo(2, 3, 4, 5, 6, 7, 8, 9)
    fmt.Println("\t The total is ", sum)
    fmt.Println("Calling Variadic Function with no arguments :: ")
    sum = foo()
    fmt.Println("\t The total is ", sum)
}

func foo(x ...int) int {
    fmt.Printf("\t Type of 'x' :: %T\n", x)
    fmt.Println("\t Values passed :: ", x)

    sum := 0
    for index, value := range x {
        sum += value
        fmt.Println("\t for item in index position, ", index, " we are now adding, ", value, " to the total which is now ", sum)
    }
    return sum
}
