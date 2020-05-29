/**
Use the given "user" struct
1. Modify the struct to work with JSON
2. create and marshal []user JSON
*/

package main

import (
    "encoding/json"
    "fmt"
)

/*
Given Struct
type user struct {
    first string
    age int
}*/

//1...
type user struct {
    First string
    Age int
}

func main() {
    //2...
    u1 := user {
        First: "James",
        Age: 32,
    }
    u2 := user {
        First: "Moneypenny",
        Age: 27,
    }

    users := []user{u1, u2}
    fmt.Println("Users Data ::", users)

    byteSlice, err := json.Marshal(users)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("JSON data ::", string(byteSlice))

}
