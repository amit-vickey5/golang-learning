package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

func main() {
    fmt.Println("CPUs ::", runtime.NumCPU())
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())

    var counter int64
    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
            fmt.Println("Loop Counter ::", atomic.LoadInt64(&counter))
            runtime.Gosched()
            wg.Done()
        }()
        fmt.Println("Loop Goroutines ::", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Main Goroutines ::", runtime.NumGoroutine())
    fmt.Println("Counter ::", counter)
}
