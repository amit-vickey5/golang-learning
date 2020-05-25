/**
1. Create a new type vehicle:
    - the underlying type is a struct
    - the fields:
        doors
        color
2. Create two new types: truck & sedan
    - the underlying type of each of these new types is a struct
    - embed the "vehicle" type in both truck and sedan
    - give truck the field "fourWheel" which will be set to bool
    - give sedan the field "luxary" which will be set to bool
3. Using the vehicle, truck, and sedan structs:
    - using a composite literal, create a value of type truck and assign values to the fields
    - using a composite literal, create a value of type sedan and assign values to the fields
4. Print out each of these values
5. Print out a single field from each of these values
*/

package main

import (
    "fmt"
)

//1...
type vehicle struct {
    doors int
    color string
}

//2...
type truck struct {
    v vehicle
    fourWheel bool
}

type sedan struct {
    vehicle
    luxary bool
}

func main() {
    //3...
    myTruck := truck {
        v: vehicle {
            doors: 2,
            color: "black",
        },
        fourWheel: false,
    }
    mySedan := sedan {
        vehicle: vehicle {
            doors: 4,
            color: "metal-grey",
        },
        luxary: true,
    }

    //4...
    fmt.Println("Truck Details :: ", myTruck)
    fmt.Println("Sedan Details :: ", mySedan)

    //5...
    fmt.Println("Truck Doors :: ", myTruck.v.doors)
    fmt.Println("Sedan Doors :: ", mySedan.doors)
}
