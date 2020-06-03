package main

import (
    "fmt"
    "runtime"
    "sync"
)

// race condition: same variable(or memory) is accessed/manipulated by multiple goroutines
// to check if race condition is there in the program....go run -race 03-race-condition.go

func main() {
    fmt.Println("CPUs ::", runtime.NumCPU())
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())

    counter := 0
    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
                v := counter
                runtime.Gosched() // yeilds the processor, allowing other goroutines to run
                v++
                counter = v
                wg.Done()
        }()
        fmt.Println("Loop Goroutines ::", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())
    fmt.Println("Counter ::", counter)
}
