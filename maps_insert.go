package main

import (
	"fmt"
)

func main() {
	mymap := map[string]string{"A": "Apple", "B": "Boy", "C": "Cat"}
	fmt.Println(mymap["A"])
	mymap["A"] = "Android"
	fmt.Println(mymap["A"])
	fmt.Printf("\nGoing to insert new value now")
	mymap["D"] = "Dog"
	fmt.Println("\n", mymap["D"])
	fmt.Println("Going to delete A")
	delete(mymap, "B")
	fmt.Println(mymap)
}
