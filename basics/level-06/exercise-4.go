/**
1. Create a user defined struct with
    - the identifier "person"
    - the fields:
        first
        last
        age
2. attach a method to type person with
    - the identifier "speak"
    - the method should have the person say their name and age
3. create a value of type person
4. call the method from the value of that person type
*/

package main

import (
    "fmt"
)

//1...
type person struct {
    first string
    last string
    age int
}

//2...
func (p person) speak() {
    fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
    //3...
    p := person {
        first: "Amit",
        last: "Cool",
        age: 24,
    }

    //4...
    p.speak()
}
