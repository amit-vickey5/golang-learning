package main

import (
    "fmt"
)

func main() {
    // buffer size is 2 i.e. at max 2 values can be stored(or kept) in the channel
    c := make(chan int, 2)

    c <- 42
    c <- 43

    // will not work as the channel size is 2, so before putting any new value
    // atlease one value needs to be extracted from the channel by some go-routine
    // c <- 44

    fmt.Println(<-c)

    // will work now as the channel size is 2, and one value have already been extracted
    // from the channel. so there's space for one more value now
    c <- 44

    fmt.Println(<-c)
    fmt.Println(<-c)
}
