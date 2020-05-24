/**
Exercise 8
    Create a map with a key of TYPE string which is a person's "last_first" name, and a value of TYPE []string
    which stores their favourite things. Store three records in your map. Print out all of the values, along
    with their index position in the slice.
        "bond_james", "Shaken, not stirred", "Maritinis", "Women"
        "moneypenny_miss", "James Bond", "Literature", "Compute Science"
        "no_dr", "Being evil", "Ice Cream", "Sunsets"

Exercise 9
    Using the code from the previous exercise, add a record to your map. Now print the map out using the "range"
    loop

Exercise 10
    Using the code from the previous exercise, delete a record from your map. Not print the map out using the
    "range" loop
*/

package main

import (
    "fmt"
)

func main() {
    //8...
    myMap := make(map[string][]string)

    myMap["bond_james"] = []string{"Shaken, not stirred", "Maritinis", "Women"}
    myMap["moneypenny_miss"] = []string{"James Bond", "Literature", "Compute Science"}
    myMap["no_dr"] = []string{"Being evil", "Ice Cream", "Sunsets"}

    fmt.Println("Exercise 8 :: ")
    for key, value := range myMap {
        fmt.Println("\tKey :: ", key)
        for index, favThings := range value {
            fmt.Println("\t\t index :: ", index, " :: ", favThings)
        }
    }

    //9...
    myMap["cool_amit"] = []string{"Women", "Photography", "Drawing"}
    fmt.Println("Exercise 9 :: ")
    for key, value := range myMap {
        fmt.Println("\tKey :: ", key)
        for index, favThings := range value {
            fmt.Println("\t\t index :: ", index, " :: ", favThings)
        }
    }

    //10...
    delete(myMap, "no_dr")
    fmt.Println("Exercise 10 :: ")
    for key, value := range myMap {
        fmt.Println("\tKey :: ", key)
        for index, favThings := range value {
            fmt.Println("\t\t index :: ", index, " :: ", favThings)
        }
    }

}
