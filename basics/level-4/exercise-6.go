/**
Create a slice to store the names of all the states. What is the length of your slice? What is
the capacity? Print out all the values along with their index position in the slice, without
using the range clause. Here is a list of all the states:

Andhra Pradesh, Arunachal Pradesh, Assam, Bihar, Chhattisgarh, Goa, Gujarat, Haryana, Himachal Pradesh,
Jharkhand, Karnataka, Kerala, Madhya Pradesh, Maharashtra, Manipur, Meghalaya, Mizoram, Nagaland,
Odisha, Punjab, Rajasthan, Sikkim, Tamil Nadu, Telangana, Tripura, Uttarakhand, Uttar Pradesh, West Bengal
*/

package main

import (
    "fmt"
)

func main() {
    states := make([]string, 28, 28)
    states = []string{
        "Andhra Pradesh",
        "Arunachal Pradesh",
        "Assam",
        "Bihar",
        "Chhattisgarh",
        "Goa",
        "Gujarat",
        "Haryana",
        "Himachal Pradesh",
        "Jharkhand",
        "Karnataka",
        "Kerala",
        "Madhya Pradesh",
        "Maharashtra",
        "Manipur",
        "Meghalaya",
        "Mizoram",
        "Nagaland",
        "Odisha",
        "Punjab",
        "Rajasthan",
        "Sikkim",
        "Tamil Nadu",
        "Telangana",
        "Tripura",
        "Uttarakhand",
        "Uttar Pradesh",
        "West Bengal",
    }

    fmt.Println("Length of slice :: ", len(states))
    fmt.Println("Capacity of slice :: ", cap(states))

    fmt.Println("List of all the states ::")
    for i := 0; i < len(states); i++ {
        fmt.Printf("\tIndex :: %v :: Value :: %v\n", i, states[i])
    }
}
