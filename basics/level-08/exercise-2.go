/**
Use the given JSON string and unmarshal it into a struct
*/

package main

import (
    jsonPkg "encoding/json"
    "fmt"
)

type person struct {
    First string
    Last string
    Age int
    Sayings []string
}

func main() {
    json := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
    fmt.Println("JSON Data ::", json)

    var people []person
    byteSlice := []byte(json)
    err := jsonPkg.Unmarshal(byteSlice, &people)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("People Data ::")
    for _, v := range people {
        fmt.Println("\tName: ", v.First, v.Last)
        fmt.Println("\tAge: ", v.Age)
        fmt.Println("\tSayings: ")
        for _, s := range v.Sayings {
            fmt.Println("\t\t", s)
        }
    }

}
