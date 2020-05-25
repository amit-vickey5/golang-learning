package main

import (
    "fmt"
)

func main() {
    a := incrementor()
    b := incrementor()
    fmt.Println("incrementor 'a' :: ")
    fmt.Println("\t", a())
    fmt.Println("\t", a())
    fmt.Println("\t", a())
    fmt.Println("\t", a())
    fmt.Println("\t", a())
    fmt.Println("incrementor 'b' :: ")
    fmt.Println("\t", b())
    fmt.Println("\t", b())
    fmt.Println("\t", b())
}

func incrementor() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}
