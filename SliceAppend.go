package main

import (
	"fmt"
)

func main() {
	arun := make([]int, 1, 4)
	fmt.Printf("Size of arun is %d. \nCapacity of arun is %d\n", len(arun), cap(arun))
	fmt.Println("Going to start appending now")
	for i := 1; i <= 10; i++ {
		arun = append(arun, i)
		fmt.Printf("Length of arun is %d and capacity is %d\n", len(arun), cap(arun))
	}
	fmt.Println(arun)

}
