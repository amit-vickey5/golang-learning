/**
Use the "defer" keyword to show that a deferred func runs after the func containing it exits.
*/

package main

import (
    "fmt"
)

func main() {
    fmt.Println("main()...1")
    defer deferredFunc()
    fmt.Println("main()...2")
}

func deferredFunc() {
    defer func() {
        fmt.Println("anonymous deferred function")
    }()
    fmt.Println("deferredFunc")
}
