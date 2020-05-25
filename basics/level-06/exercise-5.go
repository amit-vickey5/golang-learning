/**
1. create a type square
2. create a type circle
3. attach a method to each that calculates area and returns it
4. create a type shape which defines an interface as anything which has the area method
5. create a func info which takes shape and then prints the area
6. create a value of type square
7. create a value of type circle
8. use func info to print the area of the square
9. use func info to print the area of the circle
*/

package main

import (
    "fmt"
    "math"
)

//1...
type square struct {
    side float64
}

//2...
type circle struct {
    radius float64
}

//3...
func (s square) area() float64 {
    return s.side * s.side;
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

//4...
type shape interface {
    area() float64
}

//5...
func info(s shape) {
    switch s.(type) {
    case square:
        fmt.Printf("Area of square of side %v :: %v\n", s.(square).side, s.area())
    case circle:
        fmt.Printf("Area of circle of radius %v :: %v\n", s.(circle).radius, s.area())
    }
}

func main() {
    //6...
    s := square {
        side: 10,
    }

    //7...
    c := circle {
        radius: 10,
    }

    //8...
    info(s)

    //9...
    info(c)
}
