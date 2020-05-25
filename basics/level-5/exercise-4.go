/**
Create and use an anonymous struct
*/

package main

import (
    "fmt"
)

func main() {
     s := struct {
         first string
         friends map[string]int
         favDrinks []string
     }{
         first: "Amit",
         friends: map[string]int {
             "A": 555,
             "B": 666,
             "C": 777,
         },
         favDrinks: []string {
             "Wiskey",
             "Beer",
             "Hot Chocolate",
         },
     }

     fmt.Println("First Name :: ", s.first)
     fmt.Println("Friends :: ")
     for key, value := range s.friends {
         fmt.Println("\t", key, " : ", value)
     }
     fmt.Println("Favourite Drinks :: ")
     for _, value := range s.favDrinks {
         fmt.Println("\t", value)
     }
}
