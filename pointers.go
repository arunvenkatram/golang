package main
import (
	"fmt"
	"reflect"
)
func main() {
	name := "arun"
	lastname := "venkatram"
	age := 26 
	ptr := &age

	fmt.Println("Name is", name, lastname)
	fmt.Println("Type of name is ", reflect.TypeOf(name))
	fmt.Println("Memory address of age is ", &age)
	fmt.Println("Memory address of age is ", ptr , "and the content is ", *ptr)
	fmt.Println("Entering function 1 now")

	changeage(age)
	fmt.Println("After processig function1, value of age is ", age)

	fmt.Println("\nEntering fuction 2 now")
	chgrealage(&age)
	fmt.Println("Value of age after function 2 is", age)
}
func changeage(getage int) int {
	fmt.Println("Memory address of getage is ", &getage)
	getage = 30
	fmt.Println("New age is",getage)
	return getage
}
func chgrealage(addrage *int) int {
	fmt.Println("Will be trying to change address again")
	*addrage = 40
	fmt.Println("New age is", *addrage)
	return *addrage
}