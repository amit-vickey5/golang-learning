/**
Exercise 1
    Create your own type "person" which will have an underlying type of "struct" so that it can store
    the following data:
        - first name
        - last name
        - favourite ice cream flavours
    Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice.

Exercise 2
    Take the code from the previous exercise, then store the values of type person in a map with the key
    of last name. Access each value in the map. Print out the values ranaging over the slice.
*/

package main

import (
    "fmt"
)

type person struct {
    firstName string
    lastName string
    favIceCreamFlavours []string
}

func main() {
    //1...
    fmt.Println("Exercise 1 :: ")
    p1 := person {
        firstName: "Amit",
        lastName: "Cool",
        favIceCreamFlavours: []string{
            "Chocolate",
            "Strawberry",
        },
    }

    p2 := person {
        firstName: "Vickey",
        lastName: "Hot",
        favIceCreamFlavours: []string{
            "Vanilla",
            "Litchi",
            "Banana",
            "Mango",
        },
    }

    fmt.Println("\t First Person :: ")
    fmt.Println("\t\t Name :: ", p1.firstName, p1.lastName)
    fmt.Print("\t\t Favourite Ice Cream Flavours :: ")
    for _, v := range p1.favIceCreamFlavours {
        fmt.Print(v, ", ")
    }

    fmt.Println()

    fmt.Println("\t Second Person :: ")
    fmt.Println("\t\t Name :: ", p2.firstName, p2.lastName)
    fmt.Print("\t\t Favourite Ice Cream Flavours :: ")
    for _, v := range p2.favIceCreamFlavours {
        fmt.Print(v, ", ")
    }

    fmt.Println()

    //2...
    fmt.Println("Exercise 2 :: ")
    personMap := make(map[string]person)
    personMap[p1.lastName] = p1
    personMap[p2.lastName] = p2
    for key, value := range personMap {
        fmt.Println("\t Key :: ", key)
        fmt.Println("\t\t Name :: ", value.firstName, value.lastName)
        fmt.Print("\t\t Favourite Ice Cream Flavours :: ")
        for _, v := range value.favIceCreamFlavours {
            fmt.Print(v, ", ")
        }
        fmt.Println()
    }
}
