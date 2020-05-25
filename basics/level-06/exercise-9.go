/**
A "callback" is when we pass a func into a func as an argument
for this exercise,
    pass a func into a func as an argument
*/

package main

import (
    "fmt"
)

func main() {
    fmt.Println("main()...1")
    bar(foo)
    fmt.Println("main()...2")
}

func foo() {
    fmt.Println("Inside foo()")
}

func bar(f func()) {
    fmt.Println("Inside bar(). calling passed in func")
    f()
}
