/**
in addition to the main goroutine, launch two additional goroutines
    - each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("Main Routine Start")

    wg.Add(1)
    go routine1()

    wg.Add(1)
    go routine2()

    fmt.Println("Main Routine End")
    wg.Wait()
}

func routine1() {
    fmt.Println("Routine 1")
    wg.Done()
}

func routine2() {
    fmt.Println("Routine 2")
    wg.Done()
}
