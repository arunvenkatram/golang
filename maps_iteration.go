package main

import (
	"fmt"
)

//while printing, starting offset will always differ.
func main() {
	testmap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for key, value := range testmap {
		fmt.Printf("\nKey is %v, Value is %v", key, value)
	}
	fmt.Println("\n", testmap)
}
