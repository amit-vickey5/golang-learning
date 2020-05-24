package main

import (
    "fmt"
)

type BasicDetails struct {
    firstName string
    lastName string
    age int
}

type SecretAgent struct {
    details BasicDetails
    lisenseToKill bool
}

// EMBEDED STRUCT can also be created without giving a name of the embeded struct
// in this case, while populating the struct, use the type itself as the variable name
// while accessing, treat as if all the fields of embeded struct are directly a part of
// parent struct only(or can also be used with the embeded type name)
type TopSecretAgent struct {
    // anonymous field
    BasicDetails
    lisenseToKill bool
}

func main() {
    sa := SecretAgent {
        details: BasicDetails {
            firstName: "James",
            lastName: "Bond",
            age: 32,
        },
        lisenseToKill: true,
    }
    fmt.Println("Secret Agent :: ", sa)
    fmt.Println("Individual Details of Secret Agent :: ")
    fmt.Println("\t First Name :: ", sa.details.firstName)
    fmt.Println("\t Last Name :: ", sa.details.lastName)
    fmt.Println("\t Age :: ", sa.details.age)
    fmt.Println("\t Lisense To Kill :: ", sa.lisenseToKill)

    fmt.Println()

    tsa := TopSecretAgent {
        BasicDetails: BasicDetails {
            firstName: "James",
            lastName: "Bond",
            age: 32,
        },
        lisenseToKill: true,
    }
    fmt.Println("Top Secret Agent :: ", tsa)
    fmt.Println("Individual Details of Top Secret Agent :: ")
    fmt.Println("\t First Name :: ", tsa.firstName)
    fmt.Println("\t Last Name :: ", tsa.lastName)
    fmt.Println("\t Age :: ", tsa.BasicDetails.age)
    fmt.Println("\t Lisense To Kill :: ", tsa.lisenseToKill)
}
