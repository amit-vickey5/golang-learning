package main

import (
    "fmt"
)

// Calback means passing a function as a parameter

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("Numbers ::", arr)
    total := sum(arr...)
    fmt.Println("Sum of all numbers ::", total)
    evenTotal := evenSum(sum, arr...)
    fmt.Println("Sum of even numbers ::", evenTotal)
    oddTotal := evenSum(sum, arr...)
    fmt.Println("Sum of odd numbers ::", oddTotal)
}

func sum(xi ...int) int {
    total := 0
    for _, v := range xi {
        total += v
    }
    return total
}

func evenSum(sumFunction func(xi ...int) int, arr ...int) int {
    var evenArr []int
    for _, v := range arr {
        if v % 2 == 0 {
            evenArr = append(evenArr, v)
        }
    }
    total := sumFunction(evenArr...)
    return total
}

func oddSum(sumFunction func(xi ...int) int, arr ...int) int {
    var oddArr []int
    for _, v := range arr {
        if v % 2 != 0 {
            oddArr = append(oddArr, v)
        }
    }
    return sumFunction(oddArr...)
}
