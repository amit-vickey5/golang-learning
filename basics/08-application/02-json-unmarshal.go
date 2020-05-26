package main

import (
    "encoding/json"
    "fmt"
)

type person struct {
    First string
    Last string
    Age int
}

func main() {
    jsonString := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

    fmt.Println("JSON ::", jsonString)

    byteSlice := []byte(jsonString)

    var people []person

    err := json.Unmarshal(byteSlice, &people)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("DATA ::", people)

}
