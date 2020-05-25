package main

import (
    "fmt"
)


type person struct {
    first string
    last string
}

func (p person) speak() {
    fmt.Println("person.speak()")
}

func (p person) walk() {
    fmt.Println("person.walk()")
}

func (p person) run() {
    fmt.Println("person.run()")
}


type secretAgent struct {
    person
    ltk bool
}

func (sa secretAgent) speak() {
    fmt.Println("secretAgent.speak()")
}

func (sa secretAgent) walk() {
    fmt.Println("secretAgent.walk()")
}


type human interface {
    speak()
    walk()
    // this means that any type that has the method speak() and walk() is also of type human interface
}


func bar(h human) {
    fmt.Println("bar() called with parameter ::", h)
    switch h.(type) {
    case person:
        fmt.Println("\t passed parameter is 'person' type")
        fmt.Println("\t Details of Person :: Name ::", h.(person).first, h.(person).last)
        // this way to convert the interface type back to the other type is called assertion
    case secretAgent:
        fmt.Println("\t passed parameter is 'secretAgent' type")
        fmt.Println("\t Details of SecretAgent :: Name ::", h.(secretAgent).first, h.(secretAgent).last, ":: Lisence to Kill ::", h.(secretAgent).ltk)
    }
}


func main() {
    p := person {
        first: "Vickey",
        last: "Hot",
    }
    sa := secretAgent {
        person: person {
            first: "Amit",
            last: "Cool",
        },
        ltk: true,
    }
    // so, 'p' is of type 'person' and 'human' because 'person' implements all methods of 'human'
    // same goes with 'sa'. 'sa' is of both type 'secretAgent' and 'human'
    p.speak()
    p.walk()
    p.run()
    sa.speak()
    sa.walk()
    bar(p)
    bar(sa)
}
