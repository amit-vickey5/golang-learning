package main

import (
    "fmt"
)

func main() {
    p1 := struct {
        firstName string
        lastName string
    }{
        firstName: "James",
        lastName: "Bond",
    }
    fmt.Println("p1 :: ", p1)
    fmt.Println("Individual Details :: ")
    fmt.Println("\t First Name :: ", p1.firstName)
    fmt.Println("\t Last Name :: ", p1.lastName)
}
