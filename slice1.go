package main

import (
	"fmt"
)

func main() {
	//Declare a slice called arun
	arun := make([]string, 5, 10)
	venkat := []string{"apple", "orange", "grapes"}
	fmt.Printf("Length of arun is: %d. \nCapacity of arun is %d.", len(arun), cap(arun))
	fmt.Printf("\nLength of venkat is %d \nCapacity is venkat is %d", len(venkat), cap(venkat))
	fmt.Println(arun[2])
	fmt.Println(venkat[2])
}
