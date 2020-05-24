/**
Exercise 2
    1. Using a COMPOSITE LITERAL:
        a. Create a SLICE which holds 10 VALUES of TYPE INT
        b. assign VALUES to each index position
    2. Range over the slice and print the values out
    3. Using format printing
        a. print out the TYPE of the slice

Exercise 3
    1. Using the code from previous exercise, use SLICING to create the following new slices
    which are then printed:
        a. [42, 43, 44, 45, 46]
        b. [47, 48, 49, 50, 51]
        c. [44, 45, 46, 47, 48]
        d. [43, 44, 45, 46, 47]
*/

package main

import (
    "fmt"
)

func main() {
    //2-1(a)...2-1(b)...
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

    //2-2...
    for index, value := range x {
        fmt.Printf("Index :: %v :: Value :: %v\n", index, value)
    }

    //2-3...
    fmt.Printf("Type of slice :: %T\n", x)

    //3-1(a)...
    fmt.Println("[42, 43, 44, 45, 46] :: ", x[:5])

    //3-1(b)...
    fmt.Println("[47, 48, 49, 50, 51] :: ", x[5:])

    //3-1(c)...
    fmt.Println("[44, 45, 46, 47, 48] :: ", x[2:7])

    //3-1(d)...
    fmt.Println("[43, 44, 45, 46, 47] :: ", x[1:6])

}
