package main

import (
    "fmt"
)

func main() {
    s := foo()
    fmt.Printf("Type of value returned by foo() :: %T\n", s)
    fmt.Println("String returned by foo() ::", s)
    f := bar()
    fmt.Printf("Type of value returned by bar() :: %T\n", f)
    x := f()
    fmt.Println("Value returned by the function returned by bar() :: ", x)
}

func foo() string {
    return "Hello from foo()"
}

// bar() returns a function that returns an integer
func bar() func() int {
    return func() int {
        return 451
    }
}
