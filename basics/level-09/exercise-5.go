/**
Fix the race condition you created in exercise #3 by using package atomic
*/

package main

import (
    "fmt"
    // "runtime"
    "sync"
    "sync/atomic"
)

func main()  {
    var wg sync.WaitGroup
    var counter int64
    gs := 100

    wg.Add(gs)
    for i := 0; i < gs; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Counter ::", counter)
}
