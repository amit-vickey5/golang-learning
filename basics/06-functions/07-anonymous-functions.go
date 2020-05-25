package main

import (
    "fmt"
)

func main() {
    fmt.Println("Shoutout from main()")
    func() {
        fmt.Println("Shoutout from anonymous function(no parameters)")
    }()
    func(x string) {
        fmt.Println("Shoutout from anonymous function with parameter :: ", x)
    }("Hello")
}
