/**
Create a slice of a slice of string. Store the following data in the multi-dimensional slice:
    - "James", "Bond", "Shaken, not stirred"
    - "Miss", "Moneypenny", "Hellooooooo, James."
Range over the records, then range ove the data in each record
*/

package main

import (
    "fmt"
)

func main() {
    slice1 := []string{"James", "Bond", "Shaken, not stirred"}
    slice2 := []string{"Miss", "Moneypenny", "Hellooooooo, James."}

    multiSlice := [][]string {slice1, slice2}

    for index, value := range multiSlice {
        fmt.Println("Record :: ", index)
        for i, v := range value {
            fmt.Println("\tValue ", i, " :: ", v)
        }
    }
}
