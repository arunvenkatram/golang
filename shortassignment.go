package main
import (
	"fmt"
	"reflect"
)
func main() {
	name := "arun"
	lastname := "venkatram"
	age := 26 //this is short hand assignment. If declatred, it has to be used, else will throw error.

	fmt.Println("Name is", name, lastname)
	fmt.Println("Type of name is ", reflect.TypeOf(name))
}