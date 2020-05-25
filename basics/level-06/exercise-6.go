/**
Build and use an anonymous func
*/

package main

import (
    "fmt"
)

func main() {
    fmt.Println("main()...1")
    func() {
        fmt.Println("anonymous func")
    }()
    fmt.Println("main()...2")
}
