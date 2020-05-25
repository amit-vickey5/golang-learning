package main

import (
    "fmt"
)

type person struct {
    first string
    last string
}

type secretAgent struct {
    person
    ltk bool
}

// syntax :: func (r receiver) identifier(parameters) (returns)  { ... }
// when any function have a receiver of type T, any value(object) of that type T have access to that method
// in other words, when there is any receiver, the function will be attached to that receiver type, and then
// the function can be called using the dot(.) operator
func (s secretAgent) speak() {
    fmt.Println("I am ", s.first, s.last)
}

func main() {
    sa1 := secretAgent {
        person: person {
            first: "James",
            last: "Bond",
        },
        ltk: true,
    }
    sa2 := secretAgent {
        person: person {
            first: "Miss",
            last: "MoneyPenny",
        },
        ltk: false,
    }
    fmt.Print("sa1.speak() :: ")
    sa1.speak()
    fmt.Print("sa2.speak() :: ")
    sa2.speak()
}
