package main

import (
    "fmt"
)

// channels block i.e. any send(or receive) in the channel blocks the further execution of
// go-routine until another go-routing is receiving(or sending) data from the channel
// like here, main1() will not work, as on channel, 42 is being written(send), but there's
// no go-routine pulling data from the channel, so deadlock

func main1() {
    // create a channel of type int
    // means this channel can be used to transfer single int value between go-routines
    c := make(chan int)

    c <- 42 // putting value in the channel

    fmt.Println(<-c) // extracting value from channel

}

// in main(), there are two go-routines. One is putting data in the channel, and another is reading
// from the channel
func main() {
    c := make(chan int)

    go func() {
        c <- 42
    }()

    fmt.Println(<-c)

}
