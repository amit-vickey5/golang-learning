package main

import "fmt"

// a SLICE allows you to group together VALUES of the same TYPE

func main() {
	// DECLARE COMPOSITE LITERAL :: x := type{values}
	// Create SLICE using COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}

	fmt.Println("Slice :: ", x)
	fmt.Println("Length of slice :: ", len(x))
	fmt.Println("Capacity of slice :: ", cap(x))
	fmt.Println("Access elements using index :: ")
	fmt.Println("\tValue at index 0 :: ", x[0])
	fmt.Println("\tValue at index 2 :: ", x[2])

	// Loop over slice using range
	fmt.Println("Looping over slice using range :: ")
	for index, value := range x {
		fmt.Printf("\tIndex :: %v :: Value :: %v\n", index, value)
	}

	// Slicing a slice
	// Syntax :: x[lowerIndex:upperIndex]
	// lowerIndex is inclusive, but upperIndex is exclusive
	fmt.Println("Slicing a slice :: ")
	fmt.Println("\tx[:]\t:: ", x[:])
	fmt.Println("\tx[1:]\t:: ", x[1:])
	fmt.Println("\tx[1:3]\t:: ", x[1:3])

	// Append elements to a slice
	// declaration :: func append(slice []T, elements ...T) []T
	// elements ...T -> variadic parameter i.e. unlimited number of elements of type T
	fmt.Println("Append elements to a slice :: ")
	x = append(x, 77, 88, 99, 104)
	fmt.Println("\tNew Slice :: ", x)

	y := []int{10, 20, 30}
	fmt.Println("Appending a slice to another slice :: ")
	x = append(x, y...)
	fmt.Println("\tNew Slice :: ", x)

	// Deleting elements from a slice
	// Can be done by appending two slices:
	// First slice of elements before the element, Second slice of elements after the element
	fmt.Println("Deleting 2nd and 3rd element :: ")
	x = append(x[:2], x[4:]...)
	fmt.Println("\tNew Slice :: ", x)

	// Create slice using make
	// declaration :: func make(t Type, size ...IntegerType) Type
	fmt.Println("Creating slice using make ::")
	z := make([]int, 5, 7) // 5 is the length, and 7 is the capacity
	// 7 capacity means the slice can hold maximum of 7 elements
	fmt.Printf("\tLength :: %v :: Capacity :: %v :: Slice :: %v\n", len(z), cap(z), z)
	z[0] = 10
	z[4] = 20
	fmt.Printf("\tLength :: %v :: Capacity :: %v :: Slice :: %v\n", len(z), cap(z), z)
	// z[5] = 30 -> cannot do, as the length of the array is 5
	// instead, use append method
	z = append(z, 30)
	fmt.Printf("\tLength :: %v :: Capacity :: %v :: Slice :: %v\n", len(z), cap(z), z)
	z = append(z, 40)
	fmt.Printf("\tLength :: %v :: Capacity :: %v :: Slice :: %v\n", len(z), cap(z), z)
	z = append(z, 50) // now since the entire capacity is used, a new array of double the capacity will
	// be created and the underlying array will be changed to the new array
	fmt.Printf("\tLength :: %v :: Capacity :: %v :: Slice :: %v\n", len(z), cap(z), z)

	// Multi-Dimensional Slice
	slice1 := []string{"A", "B", "C"}
	slice2 := []string{"D", "E"}
	multiSlice := [][]string{slice1, slice2}
	fmt.Println("Multi-Dimensional Slice :: ", multiSlice)

}
