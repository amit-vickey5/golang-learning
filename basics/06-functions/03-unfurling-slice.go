package main

import (
    "fmt"
)

func main() {
    xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("Elemets :: ", xi)
    sum := calculateSum(xi...)
    fmt.Println("Sum :: ", sum)
}

func calculateSum(x ...int) int {
    sum := 0
    for _, value := range x {
        sum += value
    }
    return sum
}
