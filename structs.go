package main

import (
	"fmt"
)

func main() {
	type fruits struct {
		colour string
		taste  string
		count  int
	}
	//var apple fruits //Variable called apple declaraion with zero values
	//orange := new(fruits) //This is also declation of variable called orange. This one gives us a pointer.
	apple := fruits{
		colour: "red",
		taste:  "sweet",
		count:  2,
	}
	fmt.Println(apple)
	fmt.Println("Taste of apple is", apple.taste)
}
