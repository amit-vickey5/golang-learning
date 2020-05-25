package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello from main()...1")
    f1 := func() {
        fmt.Println("Shoutout from func expression(no parameters)")
    }
    f2 := func(x string) {
        fmt.Println("Shoutout from func expression with parameter :: ", x)
    }
    f1()
    fmt.Println("Hello from main()...2")
    f2("Hello")
}
