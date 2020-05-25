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

func (sa secretAgent) populateDetails() secretAgent {
    s := secretAgent {
        person: person {
            first: "Amit",
            last: "Cool",
        },
        ltk: true,
    }
    return s
}

func (sa secretAgent) printDetails() {
    fmt.Println("printDetails :: ", sa.first, " ", sa.last, " ", sa.ltk)
}

func main() {

    sa := secretAgent{}
    fmt.Println("Main 1 :: ", sa.first, " ", sa.last, " ", sa.ltk)
    sa = sa.populateDetails()
    fmt.Println("Main 2 :: ", sa.first, " ", sa.last, " ", sa.ltk)
    sa.printDetails()

}
