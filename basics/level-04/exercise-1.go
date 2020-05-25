/**
1. Using a COMPOSITE LITERAL:
    a. Create an ARRAY which holds 5 VALUES of TYPE INT
    b. assign VALUES to each index position
2. Range over the array and print the values out
3. Using format printing
    a. print out the TYPE of the arrray
*/

package main

import (
    "fmt"
)

func main() {
    //1(a)...1(b)...
    x := [5]int{10, 20, 30, 40, 50}

    //2...
    for index, value := range x {
        fmt.Printf("Index :: %v :: Value :: %v\n", index, value)
    }

    //3...
    fmt.Printf("Type of array :: %T\n", x)
}
