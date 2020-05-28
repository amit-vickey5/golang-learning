package main

import (
    "fmt"
    "sort"
)

func main() {
    xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
    xs := []string{"James", "Q", "M", "MoneyPenny", "Dr.No"}

    fmt.Println("Unsorted Integers ::", xi)
    sort.Ints(xi)
    fmt.Println("Sorted Integers   ::", xi)

    fmt.Println("Unsorted Strings  ::", xs)
    sort.Strings(xs)
    fmt.Println("Sorted Strings    ::", xs)
}
