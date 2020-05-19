package main

import "fmt"

func main() {
	//Composite Literal
	m := map[string]int{
		"Amit": 1,
		"Vickey": 2,
	}
	fmt.Println(m)
	fmt.Printf("Key :: %v :: Value :: %v\n", "Amit", m["Amit"])
	// If a non-existing key is used, it will return the zero value
	fmt.Printf("Key :: %v :: Value :: %v\n", "Amit Cool", m["Amit Cool"])

	// To check if a key is existing in the map
	v, ok := m["Amit Cool"]
	fmt.Printf("Key :: %v :: Existing in map :: %v :: Value :: %v\n", "Amit Cool", ok, v)
	v, ok = m["Amit"]
	fmt.Printf("Key :: %v :: Existing in map :: %v :: Value :: %v\n", "Amit", ok, v)
	if v, ok = m["Vickey"]; ok {
		fmt.Printf("Key :: %v :: Exists in map :: Value :: %v\n", "Vickey", m["Vickey"])
	}

	// Adding a new key to map
	m["Amit Cool"] = 3
	fmt.Println(m)

	// Access using range
	fmt.Println("Map Values :: ")
	for key, value := range m {
		fmt.Printf("\tKey :: %v :: Value :: %v\n", key, value)
	}

	// Delete an entry from map
	delete(m, "Vickey")
	fmt.Println("New Map Values :: ")
	for key, value := range m {
		fmt.Printf("\tKey :: %v :: Value :: %v\n", key, value)
	}
	// if you delete a non-existing key, it won't throw an error
	delete(m, "Cool")

	// to first check and then delete the value
	if v, ok = m["Amit"]; ok {
		fmt.Printf("\tKey :: %v :: Value :: %v\n", "Amit", v)
		delete(m, "Amit")
	}
	fmt.Println("Map :: ", m)
}
