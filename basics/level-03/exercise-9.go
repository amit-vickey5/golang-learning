/**
Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER "favSport"
 */

package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("favSport is FOOTBALL")
	case "cricket":
		fmt.Println("favSport is CRICKET")
	case "hockey":
		fmt.Println("favSport is HOCKEY")
	default:
		fmt.Println("DEFAULT")
	}
}
