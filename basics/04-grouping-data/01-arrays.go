package main

import "fmt"

func main() {
	var x [5]int // Declare an array of size 5 of INTEGER type
	fmt.Println("Array :: ", x)
	x[4] = 10 // Access any one of the index. Array is 0 based indexed
	fmt.Println("Modified Array :: ", x)
	fmt.Println("Length :: ", len(x)) // To get the length of the array
}
