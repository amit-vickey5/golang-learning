package main

import (
    "fmt"
)

func main() {
    fmt.Println("Bi-Directional Channel")
    bidirectionalChannel()
    fmt.Println("Send-Only Channel")
    sendOnlyChannel()
    fmt.Println("Receive-Only Channel")
    receiveOnlyChannel()
}

func bidirectionalChannel() {
    // bi-directional channel i.e. value can be sent/received from the channel
    biChan := make(chan int, 2)

    biChan <- 42
    biChan <- 43

    fmt.Println(<-biChan)
    fmt.Println(<-biChan)
}

func sendOnlyChannel() {
    // send-only channel i.e. value can only be sent from the channel
    sendOnlyChan := make(chan <- int, 2)

    sendOnlyChan <- 42
    sendOnlyChan <- 43

    // will not work as channel is send-only. so value cannot be received from it
    // fmt.Println(<-sendOnlyChan)
}

func receiveOnlyChannel() {
    // receive-only channel i.e. value can only be received from the channel
    receiveOnlyChan := make(<- chan int, 2)

    // will not work as channel is receive-only. so value cannot be sent from it
    // receiveOnlyChan <- 42

    fmt.Println(<-receiveOnlyChan)
}
