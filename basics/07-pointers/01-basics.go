package main

import (
    "fmt"
)

func main() {
    a := 10
    fmt.Println("Value of 'a' \t ::", a)
    fmt.Printf("Type of 'a' \t :: %T\n", a)
    fmt.Println("Address of 'a' \t ::", &a) // '&' gives the address
    fmt.Printf("Type of Address of 'a' :: %T\n", &a)

    b := &a
    fmt.Println("Value of 'b' \t ::", b)
    fmt.Printf("Type of 'b' \t :: %T\n", b)
    fmt.Println("Value stored at address stored in 'b' ::", *b) // '*' gives the value stored at that address
}
