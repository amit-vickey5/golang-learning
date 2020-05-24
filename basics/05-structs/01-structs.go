package main

import (
    "fmt"
)

// Create a new type "person" with underlying type as STRUCT
type person struct {
    firstName string
    lastName string
}

func main() {
    // Create a value of type "person"
    p1 := person{
        firstName: "James",
        lastName: "Bond",
    }
    p2 := person{
        firstName: "Miss",
        lastName: "MoneyPenny",
    }
    fmt.Println("p1 :: ", p1)

    // individual fields can be accessed using dot(.) operator
    fmt.Println("p2 :: ", p2.firstName, p2.lastName)
}
