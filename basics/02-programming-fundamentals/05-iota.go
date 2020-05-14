package main

import "fmt"

// Everytime "const" is called, "iota" value resets
// "iota" value starts with 0
// used when we want the value to be automatically incremented
// all the "iota" values are of type "int"
const (
	iotaA = iota
	iotaB = iota
	iotaC = iota
	iotaD = iota
	iotaE = iota
)

// "iota" can be declared this way as well
const (
	iotaX = iota + 10 //By default, the values will start with 0, but here the values will start with 10
	iotaY
	iotaZ
)

// all the constant variables declared after the first "iota" becomes an "iota"
// Here, iotaR and iotaS are iotas'
// But iotaP, iotaQ and iotaT are normal constants
// The value of "iota" starts with the count of the constant declared
// Like here, the value of iotaR will be 2(as it's the 3rd variable in the const block)
// and value of iotaS will be 3
const (
	iotaP = 123
	iotaQ = 12345
	iotaR = iota
	iotaS
	iotaT = "Amit"
)

func main() {
	fmt.Println("First set of iota :: ")
	fmt.Println("iotaA :: ", iotaA)
	fmt.Println("iotaB :: ", iotaB)
	fmt.Println("iotaC :: ", iotaC)
	fmt.Println("iotaD :: ", iotaD)
	fmt.Println("iotaE :: ", iotaE)
	fmt.Println("Second set of iota :: ")
	fmt.Println("iotaX :: ", iotaX)
	fmt.Println("iotaY :: ", iotaY)
	fmt.Println("iotaZ :: ", iotaZ)
	fmt.Println("Third set of iota :: ")
	fmt.Println("iotaP :: ", iotaP)
	fmt.Println("iotaQ :: ", iotaQ)
	fmt.Println("iotaR :: ", iotaR)
	fmt.Println("iotaS :: ", iotaS)
	fmt.Println("iotaT :: ", iotaT)
}
