/**
1. create a person struct
2. create a func called "changeMe" with a *person as a parameter
    - change the value stored at the *person address
3. create a value of type person
    - print out the value
4. in func main
    - call changeMe
*/

package main

import (
    "fmt"
)

//1...
type person struct {
    first string
    last string
}

//2...
func changeMe(p *person) {
    p.first = "Vickey"
    p.last = ""
}

func main() {
    //3...
    p := person {
        first: "Amit",
        last: "Cool",
    }
    fmt.Println("BEFORE :: Name ::", p.first, p.last)
    changeMe(&p)
    fmt.Println("AFTER  :: Name ::", p.first, p.last)
}
