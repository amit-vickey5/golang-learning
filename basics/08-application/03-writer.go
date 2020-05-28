package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    fmt.Println("Using fmt.Println()")
    fmt.Fprintln(os.Stdout, "Using fmt.Frintf()")
    io.WriteString(os.Stdout, "Using io.WriteString() \n")
}
