package main

import (
    "fmt"
)

func main() {
    x := 0
    fmt.Printf("BEFORE :: 'x' :: value :: %v :: address :: %v\n", x, &x)
    foo(&x)
    fmt.Printf("AFTER  :: 'x' :: value :: %v :: address :: %v\n", x, &x)
}

func foo(y *int) {
    fmt.Printf("BEFORE :: 'y' :: value :: %v :: value at address :: %v\n", y, *y)
    *y = 10
    fmt.Printf("AFTER  :: 'y' :: value :: %v :: value at address :: %v\n", y, *y)
}
