package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module float64
)

func main() {
	fmt.Println("Name is set to" , name)
	fmt.Println("Module is set to ", module)
	fmt.Println("Name type is ", reflect.TypeOf(name))
	fmt.Println("Module type is ", reflect.TypeOf(module))
}