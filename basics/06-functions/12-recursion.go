package main

import (
    "fmt"
)

func main() {
    fmt.Println("Factorial of 4 :: ", factorial(4))
    fmt.Println("Factorial of 5 :: ", factorial(5))
}

func factorial(n int) int {
    if n <= 1 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}
