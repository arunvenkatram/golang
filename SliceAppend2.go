package main

import (
	"fmt"
)

func main() {
	arun := []int{1, 2, 3, 4, 5}
	venkat := []int{10, 20, 30, 40}
	newslice := append(arun, venkat...)
	fmt.Println(newslice)
}
