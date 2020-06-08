package main

import (
    "fmt"
)

func main() {
    // buffer size is 2 i.e. at max 2 values can be stored(or kept) in the channel
    c := make(chan int, 2)

    c <- 42
    c <- 43

    fmt.Println(<-c)
    fmt.Println(<-c)
}
