package main

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

func main() {
    passStr := "MyPassword123"
    loginPass := "MyPassword123"

    passBytes, err := bcrypt.GenerateFromPassword([]byte(passStr), bcrypt.MinCost)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Password ::", passStr)
    fmt.Println("Encrypted Password ::", passBytes)
    fmt.Println("Login Password ::", loginPass)

    passValidationErr := bcrypt.CompareHashAndPassword(passBytes, []byte(loginPass))
    if passValidationErr != nil {
        fmt.Println("Incorrect Password")
        return
    }
    fmt.Println("Correct Password")
}
