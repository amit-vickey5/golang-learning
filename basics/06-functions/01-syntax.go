package main

import (
    "fmt"
)

// syntax :: func (r receiver) identifier(parameters) (returns)  { ... }
// arguments vs parameters :
// - we define our func with parameters (if any)
// - we call our func and pass in arguments (if any)
// everything in Go is PASS BY VALUE

func main() {
    foo()
    bar("James")
    s1 := woo("MoneyPenny")
    fmt.Println(s1)
    x, y := mouse("Ian", "Fleming")
    fmt.Println(x)
    fmt.Println(y)
}

// BASIC func
func foo() {
    fmt.Println("Hello from foo()")
}

// TAKES AN ARGUMENT
func bar(s string) {
    fmt.Println("Hello from bar(), ", s)
}

// SINGLE RETURN
func woo(st string) string {
    return fmt.Sprint("Hello from woo(), ", st)
}

// MULTIPLE RETURN
func mouse(fn string, ln string) (string, bool) {
    a := fmt.Sprint(fn, " ", ln, " says hello")
    b := true
    return a, b
}
