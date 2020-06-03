/**
Fix the race condition you created in the previous exercise by using a mutex
*/

package main

import (
    "fmt"
    // "runtime"
    "sync"
)

func main()  {
    var wg sync.WaitGroup
    var mu sync.Mutex
    counter := 0
    gs := 100

    wg.Add(gs)
    for i := 0; i < gs; i++ {
        go func() {
            mu.Lock()
            v := counter
            // runtime.Gosched()
            v++
            counter = v
            mu.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Counter ::", counter)
}
