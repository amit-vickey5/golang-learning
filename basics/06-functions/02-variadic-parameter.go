package main

import (
    "fmt"
)

func main() {
    sum := foo(2, 3, 4, 5, 6, 7, 8, 9)
    fmt.Println("The total is ", sum)
}

func foo(x ...int) int {
    fmt.Printf("Type of 'x' :: %T", x)
    fmt.Println("Values passed :: ", x)

    sum := 0
    for index, value := range x {
        sum += value
        fmt.Println("for item in index position, ", index, " we are now adding, ", value, " to the total which is now ", sum)
    }
    return sum
}
