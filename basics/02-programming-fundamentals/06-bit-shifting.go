package main

import "fmt"

const (
	_ = iota // not using the first iota value
	// kb = 1024 -> 2 ^ 10 -> 1 shifted by 10 bits
	kb = 1 << (iota * 10)
	// mb = 1024 * kb -> 2 ^ 20 -> 1 shifted by 20 bits
	mb = 1 << (iota * 10)
	// gb = 1034 * mb -> 2 ^ 30 -> 1 shifted by 30 bits
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("Normal Bit Shifting :: ")
	x := 4
	fmt.Printf("'x' Decimal Value :: %d :: Binary Value :: %b\n", x, x)

	// LEFT SHIFT bits by 1
	y := x << 1
	fmt.Printf("'y' Decimal Value :: %d :: Binary Value :: %b\n", y, y)

	fmt.Println("Using iota and bit shifting")
	fmt.Printf("'kb' Decimal Value :: %d\t :: Binary Value :: %b\n", kb, kb)
	fmt.Printf("'mb' Decimal Value :: %d\t :: Binary Value :: %b\n", mb, mb)
	fmt.Printf("'gb' Decimal Value :: %d :: Binary Value :: %b\n", gb, gb)
}
