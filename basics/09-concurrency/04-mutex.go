package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println("CPUs ::", runtime.NumCPU())
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())

    counter := 0
    const gs = 100
    var wg sync.WaitGroup
    var mu sync.Mutex
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            mu.Lock()
            v := counter
            runtime.Gosched() // yeilds the processor, allowing other goroutines to run
            v++
            counter = v
            mu.Unlock()
            wg.Done()
        }()
        fmt.Println("Loop Goroutines ::", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())
    fmt.Println("Counter ::", counter)
}
