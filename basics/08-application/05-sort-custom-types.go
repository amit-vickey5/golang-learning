package main

import (
    "fmt"
    "sort"
)

type person struct {
    first string
    age int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
    p1 := person{"James", 32}
    p2 := person{"MoneyPenny", 27}
    p3 := person{"Q", 64}
    p4 := person{"M", 56}

    people := []person{p1, p2, p3, p4}

    fmt.Println("Unsorted Data ::")
    for _, v := range people {
        fmt.Println("\t", v)
    }

    sort.Sort(byAge(people))
    fmt.Println("Data Sorted by Age ::")
    for _, v := range people {
        fmt.Println("\t", v)
    }

    sort.Sort(byName(people))
    fmt.Println("Data Sorted by Name ::")
    for _, v := range people {
        fmt.Println("\t", v)
    }
}
