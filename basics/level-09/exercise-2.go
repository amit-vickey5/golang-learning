/**
This exercise will reinforce our understanding of method sets:
create a type person struct
    - attach a method speak to type person using a pointer receiver
        *person
create a type human interface
    - to implicitly implement the interface, a human must have the speak method
create func “saySomething”
    - have it take in a human as a parameter
    - have it call the speak method
show the following in your code
    - you CAN pass a value of type *person into saySomething
    - you CANNOT pass a value of type person into saySomething
*/

package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func (p *person) speak() {
    fmt.Println("Person SPEAK()")
}

type human interface {
    speak()
}

func saySomething(h human) {
    h.speak()
}

func main() {
    p := person {
        name: "Amit Cool",
        age: 25,
    }
    saySomething(&p)
    // saySomething(p) -> will not work as speak is a pointer receiver
}
