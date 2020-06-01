package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("System Details ::")
    fmt.Println("\t OS ::", runtime.GOOS)
    fmt.Println("\t Architecture ::", runtime.GOARCH)
    fmt.Println("\t CPUs ::", runtime.NumCPU())
    fmt.Println("\t GoRoutines ::", runtime.NumGoroutine())

    // to run something in parallel, use "go" keyword
    wg.Add(1)
    go foo()
    bar()
    fmt.Println("\t GoRoutines ::", runtime.NumGoroutine())
    wg.Wait()
    fmt.Println("\t GoRoutines ::", runtime.NumGoroutine())
}

func foo() {
    for i := 0; i < 10; i++ {
        fmt.Println("foo :", i)
    }
    wg.Done()
}

func bar() {
    for i := 0; i < 10; i++ {
        fmt.Println("bar :", i)
    }
}
