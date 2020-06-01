package main

import (
    "fmt"
    "runtime"
)

/*
Concurrency VS Parallelism
Concurrency is an architecture. Basically a way of programming to support parallelism
Parallelism is when multiple programs are running parallely
So if there are multiple CPUs(or cores), A Concurrent program will achieve parallelism
but if there is only one CPU, even a Concurrent program will not run parallely
*/

func main() {
    fmt.Println("System Details ::")
    fmt.Println("\t OS ::", runtime.GOOS)
    fmt.Println("\t Architecture ::", runtime.GOARCH)
    fmt.Println("\t CPUs ::", runtime.NumCPU())
    fmt.Println("\t GoRoutines ::", runtime.NumGoroutine())

    // to run something in parallel, use "go" keyword
    go foo()
    bar()

    fmt.Println("\t GoRoutines ::", runtime.NumGoroutine())
}

func foo() {
    for i := 0; i < 10; i++ {
        fmt.Println("foo :", i)
    }
}

func bar() {
    for i := 0; i < 100; i++ {
        fmt.Println("bar :", i)
    }
}
