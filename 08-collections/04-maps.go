package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	productRanks["scribble-pad"] = 4
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("Iterating a map")
	for key, rank := range productRanks {
		fmt.Printf("productRanks[%q] : %d\n", key, rank)
	}

	// check for the existence of a key

	nonExistentKey := "mouse"
	keyToCheck := nonExistentKey

	/*
		existingKey := "pen"
		keyToCheck := existingKey
	*/
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] does not exist\n", keyToCheck)
	}

	// remove an item
	// keyToRemove := "pencil" //existing key
	keyToRemove := "mouse" // non existant key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key [%q], productRanks = %v\n", keyToRemove, productRanks)
}
