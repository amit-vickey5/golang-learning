/**
Assign a func to a variable, then call that func
*/

package main

import (
    "fmt"
)

func main() {
    fmt.Println("main()...1")
    f := func() {
        fmt.Println("another func")
    }
    fmt.Println("main()...2")
    f()
    fmt.Println("main()...3")
}
