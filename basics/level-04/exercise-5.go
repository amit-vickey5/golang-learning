/**
To DELETE from a slice, we use APPEND along with SLICING. For this exercise, follow these steps:
    1. start with this slice
        x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    2. use APPEND and SLICING to get the values here which you should then print:
        [42, 43, 44, 48, 49, 50, 51]
*/

package main

import (
    "fmt"
)

func main() {
    //1...
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println("Original Slice :: ", x)

    //2...
    x = append(x[:3], x[6:]...)
    fmt.Println("Modified Slice :: ", x)
}
